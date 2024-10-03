//go:build !test

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gitlab.com/golight/orm/utils"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/logs"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tools/cryptography"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/run"
)

func main() {
	err := godotenv.Load()
	conf := config.NewAppConf()
	logger := logs.NewLogger(conf, os.Stdout)
	if err != nil {
		logger.Fatal("error loading .env file")
	}
	uuidGen := cryptography.NewUUIDGenerator()
	uuid := uuidGen.String()

	config.Init(logger)
	app := run.NewApp(conf, logger)
	_ = app.Bootstrap()

	i := 0
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		uuidGen = cryptography.NewUUIDGenerator()
		uuid = uuidGen.String()
		err = app.Storages.ExchangeOrder.Create(context.Background(), models.ExchangeOrderDTO{
			UUID:       uuid,
			UserID:     666,
			ExchangeID: 1,
			Pair:       "azazga",
			Quantity:   decimal.NewFromInt(4),
			Price:      decimal.NewFromInt(999),
			Status:     5,
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
			DeletedAt:  types.NullTime{},
		})
		i++
		if i > 10 {
			break
		}
		if err != nil {
			fmt.Println("exchange order error", err)
		}
	}
	if err != nil {
		os.Exit(1)
	}

	y, err := app.Storages.ExchangeOrder.GetByUUID(context.Background(), uuid)
	err = app.Storages.ExchangeOrder.Delete(context.Background(), uuid)
	PrintJson(err)
	err = app.Storages.ExchangeOrder.Update(context.Background(), y)
	PrintJson(err)
	o, err := app.Storages.ExchangeOrder.GetList(context.Background(), utils.Condition{})
	PrintJson(o)

	//OrderLog
	err = app.Storages.ExchangeOrderLog.Create(context.Background(), models.ExchangeOrderLogDTO{
		UUID:       "peep",
		OrderID:    0,
		UserID:     228,
		ExchangeID: 300,
		Pair:       "beep",
		Quantity:   decimal.NewFromInt(544),
		Price:      decimal.NewFromInt(788),
		Status:     3,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  types.NullTime{},
	})
	z, err := app.Storages.ExchangeOrderLog.GetByID(context.Background(), 2)
	PrintJson(z)
	err = app.Storages.ExchangeOrderLog.Delete(context.Background(), 1)
	PrintJson(err)
	err = app.Storages.ExchangeOrderLog.Update(context.Background(), z)
	PrintJson(err)
	x, err := app.Storages.ExchangeOrderLog.GetList(context.Background(), utils.Condition{})
	PrintJson(x)

	//Servises

	//Delety

	app.Storages.ExchangeList.Create(context.Background(), models.ExchangeListDTO{
		ID:          0,
		Name:        "qWEF",
		Description: "eheheheh",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   types.NullTime{},
	})

	fmt.Println([]byte("\n"))

	exchangeListDelete := app.Storages.ExchangeList.Delete(context.Background(), 1)
	PrintJson(exchangeListDelete)

	exchangeUserDelete := app.Storages.ExchangeUserKey.Delete(context.Background(), 1)
	PrintJson(exchangeUserDelete)

	exchangeOrderDelete := app.Storages.ExchangeOrder.Delete(context.Background(), "43c84916-221a-11ed-bf48-acde48001122")
	PrintJson(exchangeOrderDelete)

	exchangeOrderLogDelete := app.Storages.ExchangeOrderLog.Delete(context.Background(), 1)
	PrintJson(exchangeOrderLogDelete)

	exchangeUserDelete = app.Storages.ExchangeUserKey.Delete(context.Background(), 1)
	PrintJson(exchangeUserDelete)

	//app.Storages.WebhookProcess.Create(context.Background(), models.WebhookProcessDTO{})

	webhookDelete := app.Storages.WebhookProcess.Delete(context.Background(), 5)
	PrintJson(webhookDelete)

	//Create Very
	ctx := context.Background()

	verifyStorage := app.Storages.Verifier
	if verifyStorage == nil {
		fmt.Println("verifyStorage не инициализирован")
		return
	}
	_ = verifyStorage.Create(ctx, "egobvab@gmail.com", "123", 1)
	emailVerify, err := verifyStorage.GetByEmail(ctx, "egobvab@gmail.com", "123")
	_ = verifyStorage.Create(ctx, "1@gmail.com", "223", 1)
	PrintJson(emailVerify)

	err = verifyStorage.Verify(ctx, 1)
	if err != nil {
		fmt.Println(err)
	}
	err = verifyStorage.VerifyEmail(ctx, "11@gmail.com", "223")
	if err != nil {
		fmt.Println(err)
	}
	dw, err := verifyStorage.GetByUserID(ctx, 1)
	PrintJson(dw)

	//Strategies

	app.Storages.Strategy.Create(ctx, models.StrategyDTO{
		ID:          1,
		ExchangeID:  1,
		Name:        "",
		Description: "",
		UUID:        "",
		Bots:        nil,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   types.NullTime{},
	})
	app.Storages.Strategy.Update(ctx, models.StrategyDTO{
		ID:          5,
		ExchangeID:  1,
		Name:        "",
		Description: "",
		UUID:        "",
		Bots:        nil,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   types.NullTime{},
	})
	q, err := app.Storages.Strategy.GetList(ctx)

	PrintJson(q)

	p, err := app.Storages.Strategy.GetByID(ctx, 1)

	PrintJson(p)

	app.Storages.Strategy.Delete(ctx, 5)

	//strsub

	app.Storages.StrategySubscribers.Create(ctx, models.StrategySubscribersDTO{
		ID:         0,
		ExchangeID: 2,
		UserID:     2,
		APIKey:     "eee",
		SecretKey:  "rrr",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  types.NullTime{},
	})
	m, err := app.Storages.StrategySubscribers.GetByID(ctx, 1)
	PrintJson(m)
	app.Storages.StrategySubscribers.Update(ctx, models.StrategySubscribersDTO{
		ID:         1,
		ExchangeID: 1,
		UserID:     1,
		APIKey:     "eee",
		SecretKey:  "rrr",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  types.NullTime{},
	})
	app.Storages.StrategySubscribers.Delete(ctx, 1)

	n, err := app.Storages.StrategySubscribers.GetList(ctx)
	PrintJson(n)

}
func PrintJson(in interface{}) {
	b, _ := json.Marshal(in)
	fmt.Println(string(b))
}

//func (e *EmailVerify) Verify(ctx context.Context, userID int) error {
//	dto, err := e.GetByUserID(ctx, userID)
//	if err != nil {
//		return err
//	}
//	dto.SetVerified(true)
//
//	doc, err := toDoc(dto)
//	if err != nil {
//		return err
//	}
//
//	_, err = e.getCollection().UpdateOne(ctx, bson.M{"id": dto.ID}, bson.M{"$set": doc})
//
//	return err
//}

//dto, err := e.GetByEmail(ctx, email, hash)
//	if err != nil {
//		return err
//	}
//	dto.SetVerified(true)
//
//	doc, err := toDoc(&dto)
//	if err != nil {
//		return err
//	}
//
//	_, err = e.getCollection().UpdateOne(ctx, bson.M{"id": dto.ID}, bson.M{"$set": doc})
//
//	return err
