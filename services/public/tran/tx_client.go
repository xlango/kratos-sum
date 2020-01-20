package tran

/*
	封装分布式事务Client
	TM：事务发起者
	RM：数据源代理
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"net"
	"time"
)

type Msg struct {
	GroupId string
	Type    string
	Command string
	TxCount int
	IsEnd   bool
}
type TxConnection struct {
	Tx      *gorm.DB
	Msg     Msg
	IsStart bool
}

//代理开启全局事务，TM发送begin
func (tx *TxConnection) Begin() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:7778")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	msg := Msg{
		GroupId: tx.Msg.GroupId,
		TxCount: tx.Msg.TxCount,
		Command: "create",
	}
	bytes, err := json.Marshal(&msg)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	_, err = conn.Write(bytes)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("全局分布式事务：%v 开启", tx.Msg.GroupId)
	return nil
}

//分布式事务Commit
func (tx *TxConnection) Commit() error {
	log.Info("代理分布式事务：%v Commit", tx.Msg.GroupId)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:7778")
	if err != nil {
		log.Error(err.Error())
		tx.Tx.Rollback()
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error(err.Error())
		tx.Tx.Rollback()
		conn.Close()
		return err
	}

	msg := Msg{
		GroupId: tx.Msg.GroupId,
		Type:    tx.Msg.Type,
		Command: "add",
		TxCount: tx.Msg.TxCount,
		IsEnd:   tx.Msg.IsEnd,
	}

	if tx.Msg.Command != "" {
		msg.Command = tx.Msg.Command
	}

	bytes, err := json.Marshal(&msg)
	if err != nil {
		log.Error(err.Error())
		tx.Tx.Rollback()
		conn.Close()
		return err
	}

	_, err = conn.Write(bytes)
	if err != nil {
		log.Error(err.Error())
		tx.Tx.Rollback()
		conn.Close()
		return err
	}

	if tx.IsStart {
		for {
			b := make([]byte, 1024)
			n, err := conn.Read(b)
			if err != nil {
				log.Error(err.Error())
				tx.Tx.Rollback()
				conn.Close()
				return err
			}
			reseiveMsg := make([]byte, n)
			reseiveMsg = b[:n]
			msg := Msg{}
			json.Unmarshal(reseiveMsg, &msg)
			if msg.Command == "commit" {
				//收到事务管理器通知提交
				fmt.Println(msg.Command)
				tx.Tx.Commit()
				conn.Close()
				return err
			} else if msg.Command == "rollback" {
				//收到事务管理器通知回滚
				fmt.Println(msg.Command)
				tx.Tx.Rollback()
				conn.Close()
				return err
			}

		}
	} else {
		go func() {
			for {
				b := make([]byte, 1024)
				n, err := conn.Read(b)
				if err != nil {
					log.Error(err.Error())
					tx.Tx.Rollback()
					conn.Close()
					return
				}
				reseiveMsg := make([]byte, n)
				reseiveMsg = b[:n]
				msg := Msg{}
				json.Unmarshal(reseiveMsg, &msg)
				if msg.Command == "commit" {
					//收到事务管理器通知提交
					fmt.Println(msg.Command)
					tx.Tx.Commit()
					conn.Close()
					return
				} else if msg.Command == "rollback" {
					//收到事务管理器通知回滚
					fmt.Println(msg.Command)
					tx.Tx.Rollback()
					conn.Close()
					return
				}

			}
		}()
	}

	return nil
}

func (tx *TxConnection) Rollback() error {
	log.Info("代理分布式事务：%v Rollback", tx.Msg.GroupId)
	return tx.Tx.Rollback().Error
}

func (tx *TxConnection) Close() {
	log.Info("代理分布式事务：%v Close", tx.Msg.GroupId)
	tx.Tx.Close()
}

//事务发起者 开启事务
func TMBegin(db *gorm.DB, isTM bool, txCount ...int) (txConnection *TxConnection, err error) {

	tx := db.Begin()

	txConnection = &TxConnection{
		Tx: tx,
	}

	//是否为事务发起者TM
	if isTM {
		//事务发起者生成事务组ID
		u4 := uuid.NewV4()
		txConnection.Msg.GroupId = u4.String()

		//分支事务总个数
		if len(txCount) == 0 {
			return nil, errors.New("事务发起者请设置分支事务数量")
		}

		if txCount[0] <= 0 {
			return nil, errors.New("分支事务数不能小于等于0")
		}
		txConnection.Msg.TxCount = txCount[0]

		//事务发起者标识
		txConnection.IsStart = true

		err = txConnection.Begin()

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	return txConnection, nil
}

func (tx *TxConnection) RMRollback(isEnd bool) {
	//tx.Msg.TxCount = count
	tx.Msg.IsEnd = isEnd
	tx.Msg.Type = "rollback"
}

func (tx *TxConnection) RMCommit(isEnd bool) {
	//tx.Msg.TxCount = count
	tx.Msg.IsEnd = isEnd
	tx.Msg.Type = "commit"
}

//事务发起者取消事务
func (tx *TxConnection) TMCancel() {
	tx.Msg.Command = "cancel"
}

//超时取消全局事务，并回滚
func timeout(tx *TxConnection) {
	time.Sleep(5 * time.Second)
	tx.TMCancel()
}
