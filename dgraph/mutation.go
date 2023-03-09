package dgraph

import (
	"context"
	"strings"
	"fmt"
	"time"
	"log"
	"github.com/dgraph-io/dgo/v2/protos/api"
	//"encoding/json"
)

type Event struct {
	evt_name string
	evt_loc string
}

func MuLineF(p string, s1 string, s2 string , fct string) string {
	if s1 != "" && s2!= "" {
		mu := ""
		if strings.HasPrefix(s2,"_:") || strings.HasPrefix(s2,"<0x") || strings.HasPrefix(s2,"*") {
		   mu= fmt.Sprintf(`%s <%s> %s %s .
   
		   `,s1,p,s2,fct)
		} else{
		   mu= fmt.Sprintf(`%s <%s> "%s" %s .
   
		   `,s1,p,s2,fct)
		}
		return mu
	}
	return ""
}

func MuLine(p string, s1 string, s2 string) string {
	return MuLineF(p,s1,s2,"")
}




func ExecMutation(m string) *api.Response{

	ctx := context.Background()

	nquads := []byte(m)
	
	mu := &api.Mutation{
		SetNquads: nquads,
		CommitNow: true,
	}
	resp, err1 := Server().Get().NewTxn().Mutate(ctx, mu)
	if err1 != nil {
		log.Println(err1)
		if strings.Contains(err1.Error(), "retry") {
			log.Println("Retry...")
			time.Sleep(5 * time.Second/1000)
			return ExecMutation(m)
		}
	}
	return resp
}







func ExecDelMutation(m string) *api.Response{

	ctx := context.Background()
    
	nquads := []byte(m)
	
	mu := &api.Mutation{
		DelNquads: nquads,
		CommitNow: true,
		
}
	resp, err1 := Server().Get().NewTxn().Mutate(ctx, mu)
	if err1 != nil {
		log.Println(err1)
		if strings.Contains(err1.Error(), "retry") {
			log.Println("Retry...")
			time.Sleep(5 * time.Second/1000)
			return ExecDelMutation(m)
		}
	}
	// sleep 5ms for the mutation to take effect
	time.Sleep(5 * time.Second/1000)
	return resp
}





func AddOrUpdateSchema(Schema string) {
	op := &api.Operation{}
	op.Schema = Schema 
	ctx := context.Background()
	if err := Server().Get().Alter(ctx, op); err != nil {
		log.Fatal(err)
	}

}


