package dgraph

import (
	"context"
	"log"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

func FetchData(query string) (*api.Response,error){
	txn := Server().Get().NewReadOnlyTxn().BestEffort()
	resp, err := txn.Query(context.Background(), query)
	if err != nil {
    	log.Println(err)
    }
	
	return resp,err 
}

func FetchDataWithVar(query string,vars map[string] string) (*api.Response,error){
	txn := Server().Get().NewReadOnlyTxn().BestEffort()
	resp, err := txn.QueryWithVars(context.Background(), query,vars)
	if err != nil {
    	log.Println(err)
    }
	
	return resp,err 
}
