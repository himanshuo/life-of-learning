habitrpg/  
├── bower.json - js package manager      
├── common - shared resources useful for multiple habatica repos.  
│   ├── audio  - mp3 and ogg files for various themes    
│   ├── browserify.js  
│   ├── css  
│   │   ├── backer.css  
│   │   ├── css.template.handlebars  
│   │   └── index.css  
│   ├── dist  
│   │   └── sprites  
│   ├── img  
│   │   ├── emoji  
│   │   ├── logo  
│   │   ├── project_files  
│   ├── index.js  
│   ├── locales - files specifying texts to use for given locals.  
│   ├── README.md  
│   └── script  
│       ├── constants.js - defines max health, level, and stat points  
│       ├── content - js to handle game elements  
│       │   ├── appearance - js to handle appearance  
│       │   ├── constants.js - item, events, classes, gear lists  
│       │   ├── faq.js  
│       │   ├── gear  
│       │   ├── index.js -  
│       │   ├── mystery-sets.js  
│       │   ├── spells.js  
│       │   └── translation.js  
│       ├── count.js  
│       ├── cron.js  
│       ├── fns  
│       ├── i18n.js  
│       ├── index.js  
│       ├── libs  
│       ├── ops - js for game operations  
│       ├── public  
│       │   ├── config.js  
│       │   ├── directives.js  
│       │   └── userServices.js  
│       └── statHelpers.js  
├── config.json.example  
├── CONTRIBUTING.md  
├── database_reports  
│   ├── count_users_who_own_specified_gear.js  
│   └── duplicate_indexes.js  
├── dist  
│   └── habitrpg-shared.css  
├── docker-compose.dev.yml  
├── docker-compose.yml  
├── Dockerfile  
├── Gruntfile.js  
├── gulpfile.js  
├── karma.conf.js  
├── keys  
│   └── google  
│       └── iap-live  
├── LICENSE  
├── migrations  
├── newrelic.js  
├── package.json  
├── Procfile  
├── README.md  
├── tasks  
│   ├── gulp-babelify.js  
│   ├── gulp-build.js  
│   ├── gulp-console.js  
│   ├── gulp-eslint.js  
│   ├── gulp-newstuff.js  
│   ├── gulp-sprites.js  
│   ├── gulp-start.js  
│   ├── gulp-tests.js  
│   ├── gulp-transifex-test.js  
│   └── taskHelper.js  
├── test  
│   ├── api  
│   │   ├── README.md  
│   │   └── v2  
│   │       ├── challenges  
│   │       │   └── GET-challenges_id.test.js  
│   │       ├── groups  
│   │       │   ├── chat  
│   │       │   │   ├── DELETE-groups_id_chat.test.js  
│   │       │   │   ├── GET-groups_id_chat.test.js  
│   │       │   │   ├── POST-groups_id_chat_id_clearflags.test.js  
│   │       │   │   ├── POST-groups_id_chat_id_flag.test.js  
│   │       │   │   ├── POST-groups_id_chat_id_like.test.js  
│   │       │   │   └── POST-groups_id_chat.test.js  
│   │       │   ├── GET-groups_id.test.js  
│   │       │   ├── GET-groups.test.js  
│   │       │   ├── POST-groups_id_invite.test.js  
│   │       │   ├── POST-groups_id_join.test.js  
│   │       │   ├── POST-groups_id_leave.test.js  
│   │       │   ├── POST-groups_id_removeMember.test.js  
│   │       │   ├── POST-groups_id.test.js  
│   │       │   └── POST-groups.test.js  
│   │       ├── register  
│   │       │   └── POST-register.test.js  
│   │       ├── status  
│   │       │   └── GET-status.test.js  
│   │       └── user  
│   │           ├── anonymized  
│   │           │   └── GET-user_anonymized.test.js  
│   │           ├── batch-update  
│   │           │   └── POST-user_batch-update.test.js  
│   │           ├── DELETE-user.test.js  
│   │           ├── GET-user_tags_id.test.js  
│   │           ├── GET-user_tags.test.js  
│   │           ├── GET-user.test.js  
│   │           ├── pushDevice  
│   │           │   └── POST-pushDevice.test.js  
│   │           ├── PUT-user.test.js  
│   │           └── tasks  
│   │               ├── DELETE-tasks_id.test.js  
│   │               ├── GET-tasks_id.test.js  
│   │               ├── GET-tasks.test.js  
│   │               ├── POST-tasks.test.js  
│   │               └── PUT-tasks_id.test.js  
│   ├── api-legacy  
│   ├── common  
│   │   ├── algos.mocha.js  
│   │   ├── count.js  
│   │   ├── dailies.js  
│   │   ├── preenTodos.test.js  
│   │   ├── shared.spells.test.js  
│   │   ├── simulations  
│   │   │   ├── autoAllocate.js  
│   │   │   └── passive_active_attrs.js  
│   │   ├── statHelpers.test.js  
│   │   ├── test_helper.js  
│   │   ├── user.fns.buy.test.js  
│   │   ├── user.fns.ultimateGear.test.js  
│   │   ├── user.fns.updateStats.test.js  
│   │   ├── user.ops.buyMysterySet.test.js  
│   │   ├── user.ops.equip.test.js  
│   │   ├── user.ops.hatch.js  
│   │   ├── user.ops.hourglassPurchase.test.js  
│   │   └── user.ops.test.js  
│   ├── content  
│   │   ├── faq.js  
│   │   ├── gear.js  
│   │   ├── mysterySets.test.js  
│   │   └── translaotr.js  
│   ├── e2e  
│   │   ├── helper.js  
│   │   ├── protractor.conf.js  
│   │   └── static-pages  
│   │       └── front-page.test.js  
│   ├── helpers  
│   │   ├── api-integration  
│   │   │   ├── api-classes.js  
│   │   │   ├── mongo.js  
│   │   │   ├── requester.js  
│   │   │   ├── translate.js  
│   │   │   └── v2  
│   │   │       ├── index.js  
│   │   │       └── object-generators.js  
│   │   ├── common.helper.js  
│   │   ├── content.helper.js  
│   │   └── globals.helper.js  
│   ├── migrations  
│   │   └── 20150605_ultimate_achievement_backfill.coffee  
│   ├── mocha.opts  
│   ├── README.md  
│   ├── server_side  
│   │   ├── analytics.test.js  
│   │   ├── controllers  
│   │   │   ├── groups.test.js  
│   │   │   └── user.test.js  
│   │   └── webhooks.test.js  
│   └── spec  
│       ├── app.js  
│       ├── chatServicesSpec.js  
│       ├── controllers  
│       │   ├── authCtrlSpec.js  
│       │   ├── autocompleteCtrlSpec.js  
│       │   ├── challengesCtrlSpec.js  
│       │   ├── chatCtrlSpec.js  
│       │   ├── copyMessageModalControllerSpec.js  
│       │   ├── filtersCtrlSpec.js  
│       │   ├── footerCtrlSpec.js  
│       │   ├── groupCtrlSpec.js  
│       │   ├── hallCtrlSpec.js  
│       │   ├── headerCtrlSpec.js  
│       │   ├── inventoryCtrlSpec.js  
│       │   ├── inviteToGroupCtrlSpec.js  
│       │   ├── menuCtrlSpec.js  
│       │   ├── notificationCtrlSpec.js  
│       │   ├── partyCtrlSpec.js  
│       │   ├── rootCtrlSpec.js  
│       │   ├── settingsCtrlSpec.js  
│       │   └── tasksCtrlSpec.js  
│       ├── directives  
│       │   ├── close-menu.directive.js  
│       │   ├── expand-menu.directive.js  
│       │   ├── focus-element.directive.js  
│       │   └── from-now.directive.spec.js  
│       ├── filters  
│       │   ├── largeRoundNumbersSpec.js  
│       │   ├── moneySpec.js  
│       │   └── taskOrderingSpec.js  
│       ├── mocks  
│       │   ├── analyticsMock.js  
│       │   └── sandbox.js  
│       ├── services  
│       │   ├── analyticsServicesSpec.js  
│       │   ├── groupServicesSpec.js  
│       │   ├── memberServicesSpec.js  
│       │   ├── notificationServicesSpec.js  
│       │   ├── questServicesSpec.js  
│       │   ├── statServicesSpec.js  
│       │   ├── taskServicesSpec.js  
│       │   └── userServicesSpec.js  
│       └── specHelper.js  
├── Vagrantfile.example  
├── VAGRANT.md  
├── vagrant_scripts  
│   ├── install_gcc.sh  
│   ├── install_mongo.sh  
│   ├── install_node.sh  
│   ├── install_sprite_dependencies.sh  
│   ├── install_test_dependencies.sh  
│   └── vagrant.sh  
└── website  
    ├── public  
    │   ├── 500.html  
    │   ├── apple-touch-icon-114-precomposed.png  
    │   ├── apple-touch-icon-144-precomposed.png  
    │   ├── apple-touch-icon-57-precomposed.png  
    │   ├── apple-touch-icon-72-precomposed.png  
    │   ├── apple-touch-icon-precomposed.png  
    │   ├── cake.png  
    │   ├── community-guidelines-images  
    │   │   ├── backCorner.png  
    │   │   ├── beingHabitican.png  
    │   │   ├── consequences.png  
    │   │   ├── contributing.png  
    │   │   ├── github.gif  
    │   │   ├── infractions.png  
    │   │   ├── intro.png  
    │   │   ├── moderators.png  
    │   │   ├── publicGuilds.png  
    │   │   ├── publicSpaces.png  
    │   │   ├── restoration.png  
    │   │   ├── staff.png  
    │   │   ├── tavern.png  
    │   │   ├── trello.png  
    │   │   └── wiki.png  
    │   ├── css  
    │   │   ├── alerts.styl  
    │   │   ├── avatar.styl  
    │   │   ├── challenges.styl  
    │   │   ├── classes.styl  
    │   │   ├── customizer.styl  
    │   │   ├── filters.styl  
    │   │   ├── footer.styl  
    │   │   ├── game-pane.styl  
    │   │   ├── global-colors.styl  
    │   │   ├── global-modules.styl  
    │   │   ├── header.styl  
    │   │   ├── helpers.styl  
    │   │   ├── index.styl  
    │   │   ├── inventory.styl  
    │   │   ├── items.styl  
    │   │   ├── menu.styl  
    │   │   ├── no-script.styl  
    │   │   ├── npcs.styl  
    │   │   ├── options.styl  
    │   │   ├── quests.styl  
    │   │   ├── README.md  
    │   │   ├── scrollbars.styl  
    │   │   ├── shared.styl  
    │   │   ├── static.styl  
    │   │   ├── tasks.styl  
    │   │   └── variables  
    │   │       └── screen-size.styl  
    │   ├── emails  
    │   │   └── images  
    │   │       ├── 10-days-recapture-v1.png  
    │   │       ├── 3-days-1-month-recapture-v1.png  
    │   │       ├── android-promo-v1.png  
    │   │       ├── iphone-promo-v1.png  
    │   │       ├── one-day-v1.png  
    │   │       ├── PROMO-Enchanted-Armoire-v1.png  
    │   │       ├── spring-2015-00-v1.png  
    │   │       ├── spring-2015-01-v1.png  
    │   │       ├── subscription-begins-time-travelers-v1.png  
    │   │       └── subscription-begins-v1.png  
    │   ├── favicon_192x192.png  
    │   ├── favicon.ico  
    │   ├── fontello  
    │   │   ├── css  
    │   │   │   ├── animation.css  
    │   │   │   ├── fontelico-codes.css  
    │   │   │   ├── fontelico.css  
    │   │   │   ├── fontelico-embedded.css  
    │   │   │   ├── fontelico-ie7-codes.css  
    │   │   │   └── fontelico-ie7.css  
    │   │   ├── demo.html  
    │   │   ├── font  
    │   │   │   ├── fontelico.eot  
    │   │   │   ├── fontelico.svg  
    │   │   │   ├── fontelico.ttf  
    │   │   │   └── fontelico.woff  
    │   │   ├── LICENSE.txt  
    │   │   └── README.txt  
    │   ├── front  
    │   │   ├── css  
    │   │   │   ├── blockScroll.css  
    │   │   │   ├── bootstrap.min.css  
    │   │   │   └── fixed-positioning.css  
    │   │   ├── fonts  
    │   │   │   ├── glyphicons-halflings-regular.eot  
    │   │   │   ├── glyphicons-halflings-regular.svg  
    │   │   │   ├── glyphicons-halflings-regular.ttf  
    │   │   │   ├── glyphicons-halflings-regular.woff  
    │   │   │   └── glyphicons-halflings-regular.woff2  
    │   │   ├── images  
    │   │   │   ├── achievement-perfect.png  
    │   │   │   ├── achievement-triadbingo.png  
    │   │   │   ├── avatar  
    │   │   │   │   ├── avatar.png  
    │   │   │   │   ├── avatarstatic.png  
    │   │   │   │   ├── hair_bangs_1_brown.png  
    │   │   │   │   ├── head_0.png  
    │   │   │   │   ├── head_warrior_3.png  
    │   │   │   │   ├── head_warrior_5.png  
    │   │   │   │   ├── shield_warrior_3.png  
    │   │   │   │   ├── shield_warrior_5.png  
    │   │   │   │   ├── skin_f5a76e.png  
    │   │   │   │   ├── slim_armor_warrior_3.png  
    │   │   │   │   ├── slim_armor_warrior_5.png  
    │   │   │   │   ├── slim_shirt_black.png  
    │   │   │   │   ├── Warrior.png  
    │   │   │   │   ├── weapon_healer_6.png  
    │   │   │   │   ├── weapon_warrior_3.png  
    │   │   │   │   └── weapon_warrior_5.png  
    │   │   │   ├── blackish_fox_by_kellllly-d7pzd46.png  
    │   │   │   ├── coding_by_phoneix_faerie.png  
    │   │   │   ├── devices.png  
    │   │   │   ├── explosion.jpg  
    │   │   │   ├── explosion.png  
    │   │   │   ├── Feeding_Time.png  
    │   │   │   ├── Guilds Sample Screen.png  
    │   │   │   ├── Habitica_banner_by_uncommoncriminal.png  
    │   │   │   ├── Habitica_map_by_uncommoncriminal.png  
    │   │   │   ├── habitrpg_pixel.png  
    │   │   │   ├── HabitRPGPromoPostCard6.png  
    │   │   │   ├── HabitRPGPromoThin.png  
    │   │   │   ├── Healer.png  
    │   │   │   ├── icon175x175.png  
    │   │   │   ├── intro.jpg  
    │   │   │   ├── intro.psd  
    │   │   │   ├── misc  
    │   │   │   │   ├── inventory_quest_scroll_harpy.png  
    │   │   │   │   ├── Pet_Food_Cake_Base.png  
    │   │   │   │   ├── rebirth_orb.png  
    │   │   │   │   ├── shop_gold.png  
    │   │   │   │   └── shop_potion.png  
    │   │   │   ├── mockup_for_habit_by_cosmic_caterpillar-d8mf5mb.png  
    │   │   │   ├── Mount_Body_Dragon-Golden.png  
    │   │   │   ├── Mount_Body_Dragon-Red.png  
    │   │   │   ├── Mount_Body_Wolf-Base.png  
    │   │   │   ├── Mount_Head_Dragon-Golden.png  
    │   │   │   ├── Mount_Head_Dragon-Red.png  
    │   │   │   ├── Mount_Head_Wolf-Base.png  
    │   │   │   ├── Mount.png  
    │   │   │   ├── party  
    │   │   │   │   ├── AnnaCosplay.png  
    │   │   │   │   ├── Ariel_cosplay.png  
    │   │   │   │   ├── Big_Daddy_(BioShock).png  
    │   │   │   │   ├── Cosplay_Daenerys_Targaryen.png  
    │   │   │   │   ├── GrimReaper.png  
    │   │   │   │   └── HomeStuckLusus.png  
    │   │   │   ├── Party-Header.png  
    │   │   │   ├── Pet-Dragon-Red.png  
    │   │   │   ├── Pet-Fox-Red.png  
    │   │   │   ├── presslogos  
    │   │   │   │   ├── Cnetlogo.png  
    │   │   │   │   ├── discover_logo.png  
    │   │   │   │   ├── Fast-Company-logo.png  
    │   │   │   │   ├── Forbes_logo.png  
    │   │   │   │   ├── GitHub_Logo.png  
    │   │   │   │   ├── ionic-logo-blog.png  
    │   │   │   │   ├── ionic-logo-horizontal-transparent.png  
    │   │   │   │   ├── kickstarter-logo.png  
    │   │   │   │   ├── landing_slack_hash_wordmark_logo.png  
    │   │   │   │   ├── lifehacker.png  
    │   │   │   │   ├── logo_webstorm.png  
    │   │   │   │   ├── makeuseof.png  
    │   │   │   │   ├── nyt-logo.png  
    │   │   │   │   ├── slack.png  
    │   │   │   │   └── trello-logo-blue.png  
    │   │   │   ├── Promo_springclasses2015.png  
    │   │   │   ├── Quest_dilatory_drag'on.png  
    │   │   │   ├── Quest_dilatory_drag'onSmall.png  
    │   │   │   ├── quest_vice3.png  
    │   │   │   ├── Rogue.png  
    │   │   │   ├── SAMPLEadventurers.png  
    │   │   │   ├── screenshot.png  
    │   │   │   ├── t_bone_fight_2_by_mortquitue-d8dtxbl.png  
    │   │   │   ├── testimonial_by_Streak.png  
    │   │   │   ├── testimonials  
    │   │   │   │   ├── 16bitFil.png  
    │   │   │   │   ├── AlexandraSo.png  
    │   │   │   │   ├── Althaire.png  
    │   │   │   │   ├── AndeeLiao.png  
    │   │   │   │   ├── autumnesquirrel.png  
    │   │   │   │   ├── Brenna.png  
    │   │   │   │   ├── Drag0nsilver.png  
    │   │   │   │   ├── Drei-M.png  
    │   │   │   │   ├── Elmi.png  
    │   │   │   │   ├── EvaGantz.png  
    │   │   │   │   ├── frabjabulous.png  
    │   │   │   │   ├── galarix.png  
    │   │   │   │   ├── gwyn.blath.png  
    │   │   │   │   ├── Helcura.png  
    │   │   │   │   ├── InfH.png  
    │   │   │   │   ├── irishfeet123.png  
    │   │   │   │   ├── Kai.png  
    │   │   │   │   ├── Kazui.png  
    │   │   │   │   ├── skysailor.png  
    │   │   │   │   ├── supermouse35.png  
    │   │   │   │   ├── tonitonirocca.png  
    │   │   │   │   └── Zelah_Meyer.png  
    │   │   │   ├── TVreward.png  
    │   │   │   ├── uses  
    │   │   │   │   ├── achievement-bkgd.png  
    │   │   │   │   ├── clipart-rosemonkeyct-meditation.png  
    │   │   │   │   ├── clipart-rosemonkeyct-meditation.psd  
    │   │   │   │   ├── clipart-rosemonkeyct-reading.png  
    │   │   │   │   ├── coding_3_by_phoneix_faerie-d7idtti.png  
    │   │   │   │   ├── coding.png  
    │   │   │   │   ├── consequences.png  
    │   │   │   │   ├── dusting-bkgd.png  
    │   │   │   │   ├── dusting_by_leephon.png  
    │   │   │   │   ├── gaining_an_achievement_by_cosmic_caterpillar-d7uyv5z.png  
    │   │   │   │   ├── meditation-bkgd.png  
    │   │   │   │   ├── publicSpaces.png  
    │   │   │   │   └── reading.png  
    │   │   │   ├── VICE_by_Baconsaur.png  
    │   │   │   ├── Warrior.png  
    │   │   │   └── Wizard.png  
    │   │   ├── js  
    │   │   │   ├── blockScroll.js  
    │   │   │   ├── bootstrap.min.js  
    │   │   │   └── skrollr.min.js  
    │   │   ├── landingv1Wireframe.jpg  
    │   │   ├── README.md  
    │   │   ├── staticstyle.css  
    │   │   └── style.css  
    │   ├── google280633b772b94345.html  
    │   ├── google8ca65b6ff3506fb8.html  
    │   ├── googlef3b1402b0e28338a.html  
    │   ├── js  
    │   │   ├── app.js  
    │   │   ├── controllers  
    │   │   │   ├── authCtrl.js  
    │   │   │   ├── autoCompleteCtrl.js  
    │   │   │   ├── challengesCtrl.js  
    │   │   │   ├── chatCtrl.js  
    │   │   │   ├── copyMessageModalCtrl.js  
    │   │   │   ├── filtersCtrl.js  
    │   │   │   ├── footerCtrl.js  
    │   │   │   ├── groupsCtrl.js  
    │   │   │   ├── guildsCtrl.js  
    │   │   │   ├── hallCtrl.js  
    │   │   │   ├── headerCtrl.js  
    │   │   │   ├── inventoryCtrl.js  
    │   │   │   ├── inviteToGroupCtrl.js  
    │   │   │   ├── memberModalCtrl.js  
    │   │   │   ├── menuCtrl.js  
    │   │   │   ├── notificationCtrl.js  
    │   │   │   ├── partyCtrl.js  
    │   │   │   ├── rootCtrl.js  
    │   │   │   ├── settingsCtrl.js  
    │   │   │   ├── tasksCtrl.js  
    │   │   │   ├── tavernCtrl.js  
    │   │   │   └── userCtrl.js  
    │   │   ├── directives  
    │   │   │   ├── close-menu.directive.js  
    │   │   │   ├── expand-menu.directive.js  
    │   │   │   ├── focus-element.directive.js  
    │   │   │   ├── from-now.directive.js  
    │   │   │   ├── habitrpg-tasks.directive.js  
    │   │   │   ├── hrpg-sort-checklist.directive.js  
    │   │   │   ├── hrpg-sort-tags.directive.js  
    │   │   │   ├── hrpg-sort-tasks.directive.js  
    │   │   │   ├── popover-html.directive.js  
    │   │   │   ├── popover-html-popup.directive.js  
    │   │   │   └── when-scrolled.directive.js  
    │   │   ├── env.js  
    │   │   ├── filters  
    │   │   │   ├── money.js  
    │   │   │   ├── roundLargeNumbers.js  
    │   │   │   ├── taskOrdering.js  
    │   │   │   └── timezoneOffsetToUtc.js  
    │   │   ├── services  
    │   │   │   ├── analyticsServices.js  
    │   │   │   ├── challengeServices.js  
    │   │   │   ├── chatServices.js  
    │   │   │   ├── groupServices.js  
    │   │   │   ├── guideServices.js  
    │   │   │   ├── memberServices.js  
    │   │   │   ├── notificationServices.js  
    │   │   │   ├── paymentServices.js  
    │   │   │   ├── questServices.js  
    │   │   │   ├── sharedServices.js  
    │   │   │   ├── socialServices.js  
    │   │   │   ├── statServices.js  
    │   │   │   └── taskServices.js  
    │   │   └── static.js  
    │   ├── logo  
    │   │   ├── habitrpg_bl.eps  
    │   │   ├── habitrpg.jpg  
    │   │   ├── HABITRPG-logo-version-1.gif  
    │   │   ├── HABITRPG logo version 1.psd  
    │   │   └── habitrpg_pixel.png  
    │   ├── manifest.json  
    │   ├── marketing  
    │   │   ├── android_iphone.png  
    │   │   ├── animals.png  
    │   │   ├── challenge.png  
    │   │   ├── devices.png  
    │   │   ├── drops.png  
    │   │   ├── education.png  
    │   │   ├── gear.png  
    │   │   ├── guild.png  
    │   │   ├── guild_small.png  
    │   │   ├── integration.png  
    │   │   ├── lefnire.png  
    │   │   ├── promos  
    │   │   │   ├── 201403_Forest_Walker.png  
    │   │   │   └── April14SAMPLE2.png  
    │   │   ├── screenshot.png  
    │   │   ├── social_competitve.png  
    │   │   └── wellness.png  
    │   ├── merch  
    │   │   ├── stickermule-logo.png  
    │   │   ├── stickermule-logo.svg  
    │   │   ├── stickermule.png  
    │   │   ├── teespring-eu-logo.png  
    │   │   ├── teespring-eu.png  
    │   │   ├── teespring-logo.png  
    │   │   ├── teespring-logo.svg  
    │   │   └── teespring.png  
    │   ├── page-loader.gif  
    │   ├── presskit  
    │   │   ├── Basi-List.png  
    │   │   ├── Battling the Ghost Stag.png  
    │   │   ├── Challenges Sample Screen.png  
    │   │   ├── Dread Drag'on of Dilatory.png  
    │   │   ├── Equipment Sample Screen.png  
    │   │   ├── Guilds Sample Screen.png  
    │   │   ├── Habitica Gryphon.png  
    │   │   ├── Habitica Logo Promo.png  
    │   │   ├── habitica.png  
    │   │   ├── HabiticaPromoThin.png  
    │   │   ├── Habitica Text Logo.png  
    │   │   ├── Laundromancer.png  
    │   │   ├── logo with text.png  
    │   │   ├── Market Sample Screen.png  
    │   │   ├── Necro-Vice.png  
    │   │   ├── presskit.zip  
    │   │   ├── SnackLess Monster.png  
    │   │   ├── Stagnant Dishes Boss.png  
    │   │   └── Tasks Page.png  
    │   └── refresh.png  
    ├── src  
    │   ├── controllers  
    │   │   ├── api-v2  
    │   │   │   ├── auth.js  
    │   │   │   ├── challenges.js  
    │   │   │   ├── coupon.js  
    │   │   │   ├── groups.js  
    │   │   │   ├── hall.js  
    │   │   │   ├── members.js  
    │   │   │   ├── unsubscription.js  
    │   │   │   └── user.js  
    │   │   ├── dataexport.js  
    │   │   ├── payments  
    │   │   │   ├── amazon.js  
    │   │   │   ├── iap.js  
    │   │   │   ├── index.js  
    │   │   │   ├── paypalBillingSetup.js  
    │   │   │   ├── paypal.js  
    │   │   │   └── stripe.js  
    │   │   └── pushNotifications.js  
    │   ├── libs  
    │   │   ├── analytics.js  
    │   │   ├── buildManifest.js  
    │   │   ├── firebase.js  
    │   │   ├── i18n.js  
    │   │   ├── logging.js  
    │   │   ├── utils.js  
    │   │   └── webhook.js  
    │   ├── middlewares  
    │   │   ├── apiThrottle.js  
    │   │   ├── cors.js  
    │   │   ├── domain.js  
    │   │   ├── errorHandler.js  
    │   │   ├── forceRefresh.js  
    │   │   ├── locals.js  
    │   │   └── redirects.js  
    │   ├── models  
    │   │   ├── challenge.js  
    │   │   ├── coupon.js  
    │   │   ├── emailUnsubscription.js  
    │   │   ├── group.js  
    │   │   ├── task.js  
    │   │   └── user.js  
    │   ├── routes  
    │   │   ├── api-v1.js  
    │   │   ├── api-v2  
    │   │   │   ├── auth.js  
    │   │   │   ├── coupon.js  
    │   │   │   ├── swagger.js  
    │   │   │   └── unsubscription.js  
    │   │   ├── dataexport.js  
    │   │   ├── pages.js  
    │   │   └── payments.js  
    │   └── server.js  
    └── views  
        ├── avatar-static.jade  
        ├── index.jade  
        ├── main  
        │   ├── filters.jade  
        │   └── index.jade  
        ├── options  
        │   ├── index.jade  
        │   ├── inventory  
        │   │   ├── drops.jade  
        │   │   ├── equipment.jade  
        │   │   ├── index.jade  
        │   │   ├── menu.jade  
        │   │   ├── mounts.jade  
        │   │   ├── pets.jade  
        │   │   ├── quests.jade  
        │   │   ├── seasonal-shop.jade  
        │   │   └── time-travelers.jade  
        │   ├── profile.jade  
        │   ├── settings.jade  
        │   └── social  
        │       ├── challenge-box.jade  
        │       ├── challenges.jade  
        │       ├── chat-box.jade  
        │       ├── chat-message.jade  
        │       ├── create-group.jade  
        │       ├── group.jade  
        │       ├── hall.jade  
        │       ├── index.jade  
        │       ├── party  
        │       │   ├── leave-party-and-join-another.jade  
        │       │   ├── loading-new-party.jade  
        │       │   ├── party-invitation.jade  
        │       │   └── start-a-party.jade  
        │       ├── party.jade  
        │       ├── quests  
        │       │   ├── bossStats.jade  
        │       │   ├── collectionStats.jade  
        │       │   ├── ianQuestInfo.jade  
        │       │   ├── index.jade  
        │       │   ├── participants.jade  
        │       │   ├── questActive.jade  
        │       │   └── questNotActive.jade  
        │       └── tavern.jade  
        ├── shared  
        │   ├── avatar  
        │   │   ├── appearance.jade  
        │   │   ├── generated_avatar.jade  
        │   │   └── index.jade  
        │   ├── footer.jade  
        │   ├── formatting-help.jade  
        │   ├── header  
        │   │   ├── header.jade  
        │   │   └── menu.jade  
        │   ├── mixins.jade  
        │   ├── modals  
        │   │   ├── achievements.jade  
        │   │   ├── amazon-payments.jade  
        │   │   ├── buy-gems.jade  
        │   │   ├── classes.jade  
        │   │   ├── death.jade  
        │   │   ├── drops.jade  
        │   │   ├── hatch-pet.jade  
        │   │   ├── index.jade  
        │   │   ├── invite-friends.jade  
        │   │   ├── level-up.jade  
        │   │   ├── limited.jade  
        │   │   ├── low-health.jade  
        │   │   ├── members.jade  
        │   │   ├── message-modal.jade  
        │   │   ├── new-stuff.jade  
        │   │   ├── quest-rewards.jade  
        │   │   ├── quests.jade  
        │   │   ├── raise-pet.jade  
        │   │   ├── rebirth.jade  
        │   │   ├── reroll.jade  
        │   │   ├── settings  
        │   │   │   └── change-day-start.jade  
        │   │   ├── settings.jade  
        │   │   ├── welcome.jade  
        │   │   └── won-challenge.jade  
        │   ├── new-stuff.jade  
        │   ├── noscript.jade  
        │   ├── profiles  
        │   │   ├── achievements.jade  
        │   │   ├── stats  
        │   │   │   ├── attributes.jade  
        │   │   │   ├── basic-stats.jade  
        │   │   │   ├── equipment.jade  
        │   │   │   └── pets-and-mounts.jade  
        │   │   └── stats.jade  
        │   └── tasks  
        │       ├── edit  
        │       │   ├── advanced_options.jade  
        │       │   ├── checklist.jade  
        │       │   ├── dailies  
        │       │   │   ├── calendar.jade  
        │       │   │   └── repeat_options.jade  
        │       │   ├── habits  
        │       │   │   └── plus_minus.jade  
        │       │   ├── index.jade  
        │       │   ├── rewards  
        │       │   │   └── pricing.jade  
        │       │   ├── tags.jade  
        │       │   ├── text_notes.jade  
        │       │   └── todos  
        │       │       └── due_date.jade  
        │       ├── index.jade  
        │       ├── meta_controls.jade  
        │       ├── task.jade  
        │       └── task_view  
        │           ├── add_new.jade  
        │           ├── graph.jade  
        │           ├── help.jade  
        │           ├── index.jade  
        │           ├── mixins.jade  
        │           ├── spells.jade  
        │           └── static_rewards.jade  
        ├── social  
        │   ├── achievement.jade  
        │   ├── hatch-pet.jade  
        │   ├── landing-page.jade  
        │   ├── layout.jade  
        │   ├── level-up.jade  
        │   ├── raise-pet.jade  
        │   ├── unlock-quest.jade  
        │   └── won-challenge.jade  
        └── static  
            ├── api.jade  
            ├── apps.jade  
            ├── clear-browser-data.jade  
            ├── community-guidelines.jade  
            ├── contact.jade  
            ├── faq.jade  
            ├── features.jade  
            ├── front.jade  
            ├── layout.jade  
            ├── login-modal.jade  
            ├── merch.jade  
            ├── new-stuff.jade  
            ├── old-news.jade  
            ├── overview.jade  
            ├── plans.jade  
            ├── press-kit.jade  
            ├── privacy.jade  
            ├── terms.jade  
            └── videos.jade  
