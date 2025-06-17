---
tags:
  - agent/all
  - system/config
  - phase/planning
  - doc/guide
  - tech/unity
  - tech/cpp
  - priority/critical
  - status/complete
references:
  - "[[01-Project-Overview]]"
  - "[[03-Human-AI Parallel Documentation System (HAPDS)]]"
  - "[[02-General Compression Rules (GenCompRules)]]"
---
🚫CRITICAL_BLOCKING_RULE_MUST_READ_FIRST🚫:ACTIVE_FILE_APIs_COMPLETELY_FORBIDDEN❌get_active_file❌update_active_file❌patch_active_file❌append_to_active_file→ALWAYS_USE✅get_vault_file("full/path")✅create_vault_file("full/path",content)✅patch_vault_file("full/path",...)✅append_to_vault_file("full/path",content)|VIOLATION_PREVENTION_CHECKLIST:Before_ANY_file_operation→1.Does_API_contain_"active"?→STOP_change_to_vault_API 2.Full_file_path_specified?→If_no_add_complete_path 3.Only_proceed_after_validation|

H1:ProjStruct_HAPDS_Impl:Proj_impl_Human-AI_Parallel_Doc_Sys_dual_fold_struct;001-Human/(00-SHARED/:Common_ref_docs,01-DESIGN-AGENT/:Design_agent_workspace,02-SERVER-AGENT/:Server_agent_workspace,03-CLIENT-AGENT/:Client_agent_workspace,04-LEAD-AGENT/:Lead_agent_workspace,98-USER-ONLY/:User-only_excl_MCP_ctx);002-AI-Context/(00-SHARED/:Comp_common_ref_docs,01-DESIGN-AGENT/:Comp_design_agent_workspace,02-SERVER-AGENT/:Comp_server_agent_workspace,03-CLIENT-AGENT/:Comp_client_agent_workspace,04-LEAD-AGENT/:Comp_lead_agent_workspace,98-USER-ONLY/:Comp_user_docs)|AgentAccess:DESIGN(Read:00-SHARED/,01-DESIGN-AGENT/,04-LEAD-AGENT/Integration-Reviews/;Resp:GDD&feature_spec&user_exp_design&game_balance&req_def);SERVER(Read:00-SHARED/,02-SERVER-AGENT/,01-DESIGN-AGENT/Feature-Specifications/;Resp:C++_server_arch&API_impl&DB_design&perf_opt&server_game_logic&sec);CLIENT(Read:00-SHARED/,03-CLIENT-AGENT/,01-DESIGN-AGENT/Feature-Specifications/,02-SERVER-AGENT/API-Specifications/;Resp:Unity_client_impl&UI/UX&server_comm&data_sync&plat_opt&user_exp);LEAD(Read:ALL_except_98-USER-ONLY/;Resp:Cross_agent_integ&conflict_resol&tech_arch_decisions&QA&proj_timeline_mgmt&risk_mitig)|DocStd_YAML_Frontmatter_REQ:MUST_maintain_proper_YAML_frontmatter_with_---_delimiters;tags:(agent/design|server|client|lead|all_req_choose_prim_resp_agent,system/combat|inventory|party|character|network|database|config_req_choose_main_focus,priority/critical|high|medium|low_req,status/planning|in-progress|complete|blocked_req);ref_MANDATORY_OBSIDIAN_WIKILINKS:[[Doc-Name]]_fmt_YAML_frontmatter_enable_graph_view&auto_process;!CRIT_NEW_RULE:Every_doc_MUST_include_ref_prop_YAML_frontmatter_Obsidian_wikilinks_related_docs|WorkflowPattern_Feature_Dev:1.DESIGN→Create_spec_01-DESIGN-AGENT/Feature-Specifications/→Ref_00-SHARED/_std&related_feature_specs→Include_proper_tags&wikilink_ref&cross_agent_impact→Output_complete_feature_req_depend 2.SERVER→Read_feature_spec+ref_depend→Create_API_design_02-SERVER-AGENT/API-Specifications/→Ref_design_spec&code_std&related_APIs→Output_server_impl_proper_doc 3.CLIENT→Read_feature_spec+API_spec+ref_depend→Create_UI_impl_03-CLIENT-AGENT/→Ref_design_req&server_APIs&UI_std→Output_complete_client_feature_impl 4.LEAD_integ_review→Review_all_outputs_steps_1-3→Check_cross_ref&depend_conflicts&std_compliance→Output_integ_review_recommendations|

❌FORBIDDEN_EXAMPLES_NEVER_DO❌:get_active_file()→WRONG,update_active_file("content")→WRONG,patch_active_file(...)→WRONG,append_to_active_file("content")→WRONG|✅CORRECT_EXAMPLES_ALWAYS_DO✅:get_vault_file("002-AI-Context/00-SHARED/file.md")→RIGHT,create_vault_file("002-AI-Context/00-SHARED/file.md","content")→RIGHT,patch_vault_file("002-AI-Context/00-SHARED/file.md",...)→RIGHT,append_to_vault_file("002-AI-Context/00-SHARED/file.md","content")→RIGHT|

FILE_PATH_DECLARATION_PATTERN:STEP1→Declare_TARGET_FILE="full/path/to/file.md";STEP2→Use_vault_API_with_TARGET_FILE_only;STEP3→Verify_no_"active"_in_API_call