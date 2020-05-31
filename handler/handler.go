package Handler

import (
	. "blockChain/BLC"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Data string
}

type TransData struct {
	From   []string `json:from`
	To     []string `json:to`
	Amount []string `json:amount`
}

func RespondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

/*
	处理交易请求
*/
func HandleTrans(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	jsData, _ := ioutil.ReadAll(r.Body) //获取post的数据
	var result TransData
	err := json.Unmarshal(jsData, &result)
	if err != nil {
		RespondWithJSON(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	//newBlock, err := chain.Add(nil)
	Send(result.From, result.To, result.Amount)
	RespondWithJSON(w, r, http.StatusOK, result)
}

func HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := r.Form.Get("data")

	chain := GetChain()
	newBlock, err := chain.Add(nil)
	if err != nil {
		RespondWithJSON(w, r, http.StatusInternalServerError, data)
		return
	}

	RespondWithJSON(w, r, http.StatusCreated, newBlock.LastBlock)

}

func HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	chain := GetBlockList()

	bytes, err := json.MarshalIndent(chain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
