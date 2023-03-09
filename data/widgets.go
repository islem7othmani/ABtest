package data

import (
	"fmt"
	"net/http"

	"github.com/fouita/abtesting-api/src/dgraph"
	"github.com/gin-gonic/gin"
)

// query to get data of the test
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
//get by filters
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
//create test
func CreateABTest(c *gin.Context) {
	type Test struct {
		TestAuthor     string `json:"TestAuthor,omitempty"`
		NameTest     string `json:"nameTest,omitempty"`
		TimeTest     string `json:"timeTest,omitempty"`
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
	m += dgraph.MuLine("testAuthor", "_:w", test.TestAuthor)
	m += dgraph.MuLine("nameTest", "_:w", test.NameTest)
	m += dgraph.MuLine("timeTest", "_:w", test.TimeTest)
	m += dgraph.MuLine("WidgetsTest1", "_:w", test.WidgetsTest1)
	m += dgraph.MuLine("WidgetsTest2", "_:w", test.WidgetsTest2)

	resp := dgraph.ExecMutation(m)
	c.JSON(http.StatusOK, resp)

}
// get a test by id
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
					testAuthor
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


// to delete a test
func DeleteABTest(c *gin.Context) {
	// Define a struct to hold the data to be deleted
	type dataToDelete struct {
		UID      string `json:"uid,omitempty"`
		NameTest string `json:"nameTest,omitempty"`
	}

	// Parse the request body into the struct
	var deleteData dataToDelete
	err := c.ShouldBindJSON(&deleteData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Construct the mutation string using MuLine
	m := dgraph.MuLine("nameTest", "<"+deleteData.UID+">", "*")
	/*mut:=fmt.Sprintf(`{
		delete {
			%s
		}
	}`,m)
	*/
	fmt.Println(m)
	// Execute deletion mutation
	resp := dgraph.ExecDelMutation(m)
	c.JSON(http.StatusOK, resp)
}

//to get all tests by one author (user)
func GetAllTests(c *gin.Context) {
	type widgetsByAuthor struct {
		id string
	}
	ByAuthor := widgetsByAuthor{
		id: c.Param("id"),
	}
	AllTests := fmt.Sprintf(`{
		q(func: eq(testAuthor, "%s")) {
		  uid
		  testAuthor
		  nameTest
		  timeTest
		  WidgetsTest1
		  widgetsTest2
		}
	  }
	  `,ByAuthor.id)

	vars := map[string]string{"$id":ByAuthor.id}
	res, _ := dgraph.FetchDataWithVar(AllTests, vars)

	c.JSON(http.StatusOK, res)

}