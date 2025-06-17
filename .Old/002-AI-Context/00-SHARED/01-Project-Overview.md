---
tags:
  - agent/all
  - phase/planning
  - doc/gdd
  - tech/unity
  - tech/cpp
  - priority/critical
  - status/complete
references:
  - "[[04-Coding-Standards]]"
  - "[[05-MCP-Usage-Guide]]"
  - "[[03-Human-AI Parallel Documentation System (HAPDS)]]"
  - "[[02-General Compression Rules (GenCompRules)]]"
---

ProjInfo:ProjName:RaidMaster;Genre:2D_Action_RPG;Platform:PC(Win/Mac/Linux);Timeline:6_months_solo_dev;DevProfile:2yrs_WPF_exp_new_game_dev|TechStack:Client(Engine:Unity_2023.x_LTS,Lang:C#,UI:Unity_UI_Toolkit,Render:URP);Server(Lang:C++_17/20,Net:Boost.Asio/libuv,DB:MySQL/PostgreSQL,Cache:Redis,Build:CMake);Common(VCS:Git,Doc:Obsidian+Claude_MCP,CI/CD:GitHub_Actions_planned)|ProjPrinciples:1.Practicality_First:Working_code>perfect_code 2.Incremental_Dev:MVP_first_feature_add_later 3.Doc_Focus:All_decisions&designs_must_documented 4.AI_Assisted_Dev:Leverage_Claude_agents_efficient_dev|AgentWorkflow:Responsibilities(DESIGN:GDD&feature_specs&user_stories,SERVER:C++_server_arch&API_design&DB_schema,CLIENT:Unity_impl&UI/UX&client_server_comm,LEAD:Integration_mgmt&tech_decisions&QA);Access_Rights(DESIGN:Work_01-DESIGN-AGENT/&read_00-SHARED/&04-LEAD-AGENT/Integration-Reviews/,SERVER:Work_02-SERVER-AGENT/&read_00-SHARED/&01-DESIGN-AGENT/Feature-Specifications/,CLIENT:Work_03-CLIENT-AGENT/&read_00-SHARED/&01-DESIGN-AGENT/Feature-Specifications/&02-SERVER-AGENT/API-Specifications/,LEAD:Full_access_all_except_98-USER-ONLY/)|HAPDS_Impl:Dual_fold_struct(001-Human/:Human_readable_doc,002-AI-Context/:Comp_AI_consumption);Sync_both_fold_maint_ident_sem_cont_diff_fmt;MCP_Ctx_Strat(AI_Agents:Use_002-AI-Context/_token_opt_ctx,Human_Work:Use_001-Human/_daily_work&edit,Sync:Both_fold_ident_sem_cont_diff_fmt);Fold_Struct(001-Human/&002-AI-Context/:00-SHARED/_common_proj_info,01-DESIGN-AGENT/_design_docs,02-SERVER-AGENT/_server_docs,03-CLIENT-AGENT/_client_docs,04-LEAD-AGENT/_integ_mgmt,98-USER-ONLY/_user_only_templates_excl_MCP_ctx)|CurrentStatus:Phase:Initial_proj_setup;Completed:Obsidian_MCP_integ&fold_struct_design;NextSteps:Game_concept_final&GDD_creation|DevGoals:Primary:Complete_playable_2D_Action_RPG_6_months;Secondary:Establish_efficient_AI_assisted_dev_workflow;Quality:Focus_core_gameplay>visual_polish;Learning:Gain_practical_game_dev_exp|DocRef:Related(00-SHARED/Document-Standards.md:Doc&tag_req,00-SHARED/04-Coding-Standards.md:Impl_std_all_agents,00-SHARED/05-MCP-Usage-Guide.md:Agent_workflow&access_patterns,01-DESIGN-AGENT/Game-Design-Documents/Main-GDD.md:Detailed_game_design);Depend(Req:None_foundation_doc,Blocks:All_other_proj_docs_depend_this_overview);CrossAgent_Impact(DESIGN:Provides_tech_constraints&proj_scope_game_design,SERVER:Defines_server_tech_stack&perf_req,CLIENT:Specifies_Unity_ver&platform_targets&client_req,LEAD:Establishes_proj_timeline&dev_methodology)