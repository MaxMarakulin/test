package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"test/database"
)

func getExchangeRate(c *gin.Context) {
	// Просто получаем все, что есть в таблице с курсами
	var records []response
	query := database.DB.Table("rates.exchange_rate er")
	query.Select("er.id, c_from.name AS from, c_to.name AS to, er.value, er.source")
	query.Joins("LEFT JOIN rates.currencies c_from ON c_from.id = er.from")
	query.Joins("LEFT JOIN rates.currencies c_to ON c_to.id = er.to")
	query.Scan(&records)

	c.JSON(http.StatusOK, &records)
}

func exchangeRate() {
	wg := &sync.WaitGroup{}
	// Ставим значение счетчика 2, что соответствует количеству горутин
	wg.Add(2)
	// Берем два источника, это центробанк России и национальный банк республики Беларусь, т.к. у них апи в свободном доступе и не требуют регистрации
	go exchangeRateCBR(wg)
	go exchangeRateNBRB(wg)
	// Эти две функции по коду почти одинаковы и если постараться, можно сделать одну общую,
	// но такая реализация - самый простой способ наглядно использовать горутины
	// Ждем, пока все горутины отработают
	wg.Wait()
}

func exchangeRateCBR(wg *sync.WaitGroup) {
	// В самом конце выполнения функции указываем, что она отработала
	defer wg.Done()

	// Так как источник центробанк России, сразу указываем конкретную валюту
	var currencyName = "RUB"
	var url = "https://www.cbr-xml-daily.ru/daily_json.js"

	bodyCBR, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err)
	}
	defer func() {
		err = bodyCBR.Body.Close()
		if err != nil {
			log.Println("body close error:", err)
		}
	}()
	var responseCBR rateCBR
	err = json.NewDecoder(bodyCBR.Body).Decode(&responseCBR)
	if err != nil {
		log.Println("Error:", err)
	}
	var records exchangeRateModel

	// Находим id валюты RUB
	records.From = getCurrencyId(currencyName)
	// Так же поступаем с валютой которая котируется по отношению к RUB
	records.To = getCurrencyId(responseCBR.Currency.BYN.CharCode)
	records.Value = responseCBR.Currency.BYN.Value
	records.Source = url
	// Обновляем курс, если он уже существует, если нет, то добавляем
	createOrUpdate(records)

}

func exchangeRateNBRB(wg *sync.WaitGroup) {
	// Почти один в один как и в exchangeRateCBR, за исключением разницы в валютах, адресе и структуре
	defer wg.Done()

	var currencyName = "BYN"
	var url = "https://api.nbrb.by/exrates/rates/456"

	bodyNBRB, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err)
	}
	defer func() {
		err = bodyNBRB.Body.Close()
		if err != nil {
			log.Println("body close error:", err)
		}
	}()
	var responseNBRB rateNBRB
	err = json.NewDecoder(bodyNBRB.Body).Decode(&responseNBRB)
	if err != nil {
		log.Println("Error:", err)
	}
	var records exchangeRateModel

	records.From = getCurrencyId(currencyName)
	records.To = getCurrencyId(responseNBRB.CurAbbreviation)
	records.Value = responseNBRB.CurOfficialRate
	records.Source = url
	createOrUpdate(records)
}

func getCurrencyId(name string) int32 {
	var id_ int32
	// Пробуем найти идентификатор валюты, если такового нет, то добавляем
	result := database.DB.Table("rates.currencies").Select("id").Where("name = ?", name).Take(&id_).RowsAffected
	if result == 0 {
		currencies := currenciesModel{Name: name}
		database.DB.Table("rates.currencies").Create(&currencies)
		id_ = currencies.Id
	}
	return id_
}

func createOrUpdate(value exchangeRateModel) {
	result := database.DB.Table("rates.exchange_rate").Where(`"to" = ? AND "from" = ?`, value.To, value.From).Updates(&value).RowsAffected
	if result == 0 {
		database.DB.Table("rates.exchange_rate").Create(&value)
	}
}
