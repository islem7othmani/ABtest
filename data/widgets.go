package data

import (
	"fmt"
	"net/http"
	"github.com/fouita/abtesting-api/src/dgraph"
	"github.com/gin-gonic/gin"
	//"github.com/graphql-go/graphql"

)
//query to get data of the test
func GetAnalytics(c *gin.Context) {
	type WidgetElements struct {
		id string
	}
	
	widgetElements := WidgetElements{
		id: c.Param("id"),
	}
	
	q := fmt.Sprintf(`{
		q(func: uid(%s)) {
			uid
			nameW
			description
			img
			v
			w_live
			cloned_at
			updated_at
			widget_views
			vcmp_action{
				ac_name
				ac_count
			}
			cln_cmp_from{
				uid
				ac_engagement
				ac_conversion
			}
			w_owner{
				uid
				u_analytics_enabled
				need_upgrade
				u_month_views
				u_reset_views_at
				u_month_views_quota
			}
			forms_{
				uid
				frm_data
				
			}
			n_forms:count(forms_)
			origin: cln_cmp_from {
				uid
				ac_engagement
				ac_conversion
			}
			actions: vcmp_action @groupby(ac_name) @filter(lt(ac_count,500)){
				num_actions: sum(ac_count)
			}
			action_summary: vcmp_action @groupby(name: ac_name) {
				num_actions: sum(ac_count)
			}
		}
	}`, widgetElements.id)

	vars := map[string]string{"$id": widgetElements.id}
	res, _ := dgraph.FetchDataWithVar(q, vars)

	c.JSON(http.StatusOK, res)
}


func GetFilter(c *gin.Context) {

	qF := fmt.Sprintf(`{
		var(func: uid(111)) {
			vcmp_action  {
			v as ac_visitor
			}
			}
			byCountry(func: uid(v)) @groupby(country:v_country) {
			count(uid)
			}
			byBrowser(func: uid(v)) @groupby(browser:v_browser) {
			count(uid)
			}
			byOs(func: uid(v)) @groupby(os:v_os) {
			count(uid)
			}
			byOrigin(func: uid(v)) @groupby(origin:v_origin) {
			count(uid)
			}
			byDevice(func: uid(v)) @groupby(device:v_device) {
			count(uid)
			}
			byReferrer(func: uid(v)) @groupby(referrer:v_referrer) {
			count(uid)
			}
			byUtmSource(func: uid(v)) @groupby(utm_source:v_utm_source) {
			count(uid)
			}
			byUtmMedium(func: uid(v)) @groupby(utm_medium:v_utm_medium) {
			count(uid)
			}
			byUtmCampaign(func: uid(v)) @groupby(utm_campaign:v_utm_campaign) {
			count(uid)
			}
			byUtmContent(func: uid(v)) @groupby(utm_content:v_utm_content) {
			count(uid)
			}
			byUtmTerm(func: uid(v)) @groupby(utm_term:v_utm_term) {
			count(uid)
			}
	}`)
     
	res, _ := dgraph.FetchData(qF)

	c.JSON(http.StatusOK, res)

}



func CreateABTest(c *gin.Context) {
	type Test struct {
		NameTest string `json:"nameTest,omitempty"`
		TimeTest string `json:"timeTest,omitempty"`
		WidgetsTest1 string `json:"WidgetsTest1,omitempty"`
		WidgetsTest2 string `json:"WidgetsTest2,omitempty"`
	}
	
	var test Test
	err := c.ShouldBindJSON(&test)
	if err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}

	m := ""
	m += dgraph.MuLine("nameTest", "_:w", test.NameTest)
	m += dgraph.MuLine("timeTest", "_:w", test.TimeTest)
	m += dgraph.MuLine("WidgetsTest1", "_:w", test.WidgetsTest1)
	m += dgraph.MuLine("WidgetsTest2", "_:w", test.WidgetsTest2)

	resp := dgraph.ExecMutation(m)
	c.JSON(http.StatusOK, resp)

	
}

func GetABTest(c *gin.Context) {
	type TestElements struct {
		id string
	}
	
	TestID := TestElements{
		id: c.Param("id"),
	}
	
	qAB := fmt.Sprintf(`{
		q(func:uid(%s)){
					uid
					nameTest
					timeTest
					WidgetsTest1
					WidgetsTest2
				}
			  
	}`, TestID.id)

	vars := map[string]string{"$id": TestID.id}
	res, _ := dgraph.FetchDataWithVar(qAB, vars)

	c.JSON(http.StatusOK, res)
}





