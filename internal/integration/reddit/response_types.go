package reddit

// Auto-generated from json
// TODO: clean up interface
type SubredditListing struct {
	Kind string `json:"kind"`
	Data struct {
		After     any    `json:"after"`
		Dist      int    `json:"dist"`
		Modhash   string `json:"modhash"`
		GeoFilter string `json:"geo_filter"`
		Children  []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc              any     `json:"approved_at_utc"`
				Subreddit                  string  `json:"subreddit"`
				Selftext                   string  `json:"selftext"`
				AuthorFullname             string  `json:"author_fullname"`
				Saved                      bool    `json:"saved"`
				ModReasonTitle             any     `json:"mod_reason_title"`
				Gilded                     int     `json:"gilded"`
				Clicked                    bool    `json:"clicked"`
				Title                      string  `json:"title"`
				LinkFlairRichtext          []any   `json:"link_flair_richtext"`
				SubredditNamePrefixed      string  `json:"subreddit_name_prefixed"`
				Hidden                     bool    `json:"hidden"`
				Pwls                       int     `json:"pwls"`
				LinkFlairCSSClass          any     `json:"link_flair_css_class"`
				Downs                      int     `json:"downs"`
				TopAwardedType             any     `json:"top_awarded_type"`
				HideScore                  bool    `json:"hide_score"`
				Name                       string  `json:"name"`
				Quarantine                 bool    `json:"quarantine"`
				LinkFlairTextColor         string  `json:"link_flair_text_color"`
				UpvoteRatio                float64 `json:"upvote_ratio"`
				AuthorFlairBackgroundColor any     `json:"author_flair_background_color"`
				SubredditType              string  `json:"subreddit_type"`
				Ups                        int     `json:"ups"`
				TotalAwardsReceived        int     `json:"total_awards_received"`
				MediaEmbed                 struct {
				} `json:"media_embed"`
				AuthorFlairTemplateID any   `json:"author_flair_template_id"`
				IsOriginalContent     bool  `json:"is_original_content"`
				UserReports           []any `json:"user_reports"`
				SecureMedia           any   `json:"secure_media"`
				IsRedditMediaDomain   bool  `json:"is_reddit_media_domain"`
				IsMeta                bool  `json:"is_meta"`
				Category              any   `json:"category"`
				SecureMediaEmbed      struct {
				} `json:"secure_media_embed"`
				LinkFlairText       any    `json:"link_flair_text"`
				CanModPost          bool   `json:"can_mod_post"`
				Score               int    `json:"score"`
				ApprovedBy          any    `json:"approved_by"`
				IsCreatedFromAdsUI  bool   `json:"is_created_from_ads_ui"`
				AuthorPremium       bool   `json:"author_premium"`
				Thumbnail           string `json:"thumbnail"`
				AuthorFlairCSSClass any    `json:"author_flair_css_class"`
				AuthorFlairRichtext []any  `json:"author_flair_richtext"`
				Gildings            struct {
				} `json:"gildings"`
				ContentCategories        any     `json:"content_categories"`
				IsSelf                   bool    `json:"is_self"`
				ModNote                  any     `json:"mod_note"`
				Created                  float64 `json:"created"`
				LinkFlairType            string  `json:"link_flair_type"`
				Wls                      int     `json:"wls"`
				RemovedByCategory        any     `json:"removed_by_category"`
				BannedBy                 any     `json:"banned_by"`
				AuthorFlairType          string  `json:"author_flair_type"`
				Domain                   string  `json:"domain"`
				AllowLiveComments        bool    `json:"allow_live_comments"`
				SelftextHTML             any     `json:"selftext_html"`
				Likes                    any     `json:"likes"`
				SuggestedSort            any     `json:"suggested_sort"`
				BannedAtUtc              any     `json:"banned_at_utc"`
				URLOverriddenByDest      string  `json:"url_overridden_by_dest"`
				ViewCount                any     `json:"view_count"`
				Archived                 bool    `json:"archived"`
				NoFollow                 bool    `json:"no_follow"`
				IsCrosspostable          bool    `json:"is_crosspostable"`
				Pinned                   bool    `json:"pinned"`
				Over18                   bool    `json:"over_18"`
				AllAwardings             []any   `json:"all_awardings"`
				Awarders                 []any   `json:"awarders"`
				MediaOnly                bool    `json:"media_only"`
				CanGild                  bool    `json:"can_gild"`
				Spoiler                  bool    `json:"spoiler"`
				Locked                   bool    `json:"locked"`
				AuthorFlairText          any     `json:"author_flair_text"`
				TreatmentTags            []any   `json:"treatment_tags"`
				Visited                  bool    `json:"visited"`
				RemovedBy                any     `json:"removed_by"`
				NumReports               any     `json:"num_reports"`
				Distinguished            any     `json:"distinguished"`
				SubredditID              string  `json:"subreddit_id"`
				AuthorIsBlocked          bool    `json:"author_is_blocked"`
				ModReasonBy              any     `json:"mod_reason_by"`
				RemovalReason            any     `json:"removal_reason"`
				LinkFlairBackgroundColor string  `json:"link_flair_background_color"`
				ID                       string  `json:"id"`
				IsRobotIndexable         bool    `json:"is_robot_indexable"`
				ReportReasons            any     `json:"report_reasons"`
				Author                   string  `json:"author"`
				DiscussionType           any     `json:"discussion_type"`
				NumComments              int     `json:"num_comments"`
				SendReplies              bool    `json:"send_replies"`
				ContestMode              bool    `json:"contest_mode"`
				ModReports               []any   `json:"mod_reports"`
				AuthorPatreonFlair       bool    `json:"author_patreon_flair"`
				AuthorFlairTextColor     any     `json:"author_flair_text_color"`
				Permalink                string  `json:"permalink"`
				Stickied                 bool    `json:"stickied"`
				URL                      string  `json:"url"`
				SubredditSubscribers     int     `json:"subreddit_subscribers"`
				CreatedUtc               float64 `json:"created_utc"`
				NumCrossposts            int     `json:"num_crossposts"`
				Media                    any     `json:"media"`
				IsVideo                  bool    `json:"is_video"`
			} `json:"data"`
		} `json:"children"`
		Before any `json:"before"`
	} `json:"data"`
}
