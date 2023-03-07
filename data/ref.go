package data

/*


func GetWidgetTrackSummary(c *gin.Context) {

	// TODO - check if user owns the vcmp!

	// search filters
	// show analytics by day: group hours
	var vcmp struct {
		Vuid     string
		Interval string
		Actions  []string
		After    string // filter after this datetime
		Device   string // filter by device from list (desktop, laptop, tablet, mobile)
		Os       string // filter by OS, search by all OS
		Location string // filter by location, all locations, group by location

	}
	if err := c.ShouldBindJSON(&vcmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !access.HasAccessFromWidget(c, vcmp.Vuid) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}
	// get summary of vcmp access (not project)
	// group by day, last month or last 7 days ? by default
	group := ""

	filter := fmt.Sprintf("gt(ac_time,%q)", vcmp.After)

	if vcmp.Interval == "Day" {
		group = "time:ac_time,name:ac_name"
	} else if vcmp.Interval == "Week" || vcmp.Interval == "Month" {
		group = "time:ac_day,name:ac_name"
	} else if vcmp.Interval == "HalfYear" || vcmp.Interval == "Year" {
		group = "time:ac_month,name:ac_name"
	}

	acFilter := ""
	acNames := "[" + strings.Join(vcmp.Actions, ",") + "]"

	if len(vcmp.Actions) > 0 {
		acFilter = fmt.Sprintf("and eq(ac_name, %s)", acNames)
		filter += acFilter
	}

	q := fmt.Sprintf(`{
		var(func: uid(%s)) {
			vcmp_action @filter(%s) {
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
		q(func: uid(%s)) {
			uid
			name
			description
			img
			v
			live: w_live
			created: cloned_at
			updated: updated_at
			n_forms: count(forms)
			actions: vcmp_action @groupby(%s) @filter(%s){
				num_actions: sum(ac_count)
			}
			action_summary: vcmp_action @groupby(name: ac_name) @filter(%s) {
				num_actions: sum(ac_count)
			}
			origin: cln_cmp_from {
				uid
				ac_engagement
				ac_conversion
			}
			views: widget_views
			forms {
				n_submissions: count(frm_data)
			}
			owner: w_owner {
				uid
				analytics_enabled: u_analytics_enabled
				upgrade: need_upgrade
				v: u_month_views
				vrt: u_reset_views_at
				vq: u_month_views_quota
			}
		}
	}`, vcmp.Vuid, filter, vcmp.Vuid, group, filter, filter)

	resp, _ := query.FetchData(q)

	c.JSON(http.StatusOK, resp)

}

*/

/*

func MyWidgets(c *gin.Context) {
	uuid := c.GetHeader("X-UID")

	var data struct {
		Tags string `json:"tags,omitempty"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := ""

	if data.Tags != "" {
		// Cannot be %q (TO FIX - TODO)
		filter = fmt.Sprintf("@filter(uid_in(w_tags, %s))", data.Tags)
	}

	ftime := fmt.Sprintf("gt(ac_time,%q)", timeLastMonth())

	q := fmt.Sprintf(`{
		q(func: uid(%s)) {
			uid
			widgets: ~w_owner (orderdesc: cloned_at) %s {
				uid
				name
				description
				url
				img
				v
				created: cloned_at
				live: w_live
				active: w_active
				id: css_key
				wv: widget_views
				origin: cln_cmp_from {
					uid
					ac_engagement
					ac_conversion
				}
				forms {
					n_submissions: count(frm_data)
				}
				action_summary: vcmp_action @groupby(name: ac_name) @filter(%s) {
					num_actions: sum(ac_count)
				}
				w_tags {
					uid
					name: w_tag_name
					color: w_tag_color
				}
			}
		}
	}`, uuid, filter, ftime)

	res, _ := query.FetchData(q)

	c.JSON(http.StatusOK, res)
}

*/


