---
tags:
  - agent/all
  - system/config
  - doc/guide
  - priority/critical
  - status/planning
references:
  - "[[03-Human-AI Parallel Documentation System (HAPDS)]]"
  - "[[02-General Compression Rules (GenCompRules)]]"
  - "[[05-MCP-Usage-Guide]]"
---

CL_HAPDS:Co_Located_Human_AI_Parallel_Doc_Sys|Purp:Maint_dual_doc_fmt_same_fold_opt_human_comp&AI_ctx_consume;Elim_dual_fold_struct_complex|Prin:Every_proj_doc=2_co_located_files(Human:filename.md_Obsidian_native,AI:filename.md.aicomp_hidden_Obsidian)|Struct:Proj_fold(filename.md+filename.md.aicomp,arch.md+arch.md.aicomp,reqs.md+reqs.md.aicomp)|FileExt:.aicomp(AI_comp_ver_hidden_Obsidian),.md(Human_readable_Obsidian_native)|ObsidianConfig:Add_ignore_pattern(*.aicomp,*.ctx,*.tok)_hide_comp_files_graph&search|APILogic:get_vault_file()→Check_filename.md.aicomp_exists→Return_comp_ver_AI_ctx:Return_human_ver_other_use|SyncRules:1.SemEquiv:Both_ver=ident_sem_info,rel,crit_detail;Only_present_fmt_diff 2.CoLocation:Comp_file_adjacent_human_file_same_fold 3.UpdateProp:Cont_chg_human_ver→Comp_ver_same_fold_upd_sync 4.VerSrcPrio:Human_ver=prim_src_cont_create&maj_upd;AI_comp_ver_gen_human_ver;Both_ver=auth_resp_use_case|Benefits:SimpleFoldStruct(Single_fold_hier_no_dual_001/002_complex),EasierMaint(Comp_file_next_source_easy_sync),CleanObsidian(Graph_view_only_show_human_files),FlexAPI(Intel_choose_ver_based_ctx_need),ReducedComplexity(No_fold_struct_mismatch_risk)|ImplReq:ObsidianSetup(Config_ignore_*.aicomp_files),APIUpgrade(Modify_get_vault_file()_check_comp_ver_first),FileNaming(Use_.aicomp_ext_AI_comp_ver),MaintWorkflow(Upd_human_first→Gen_comp_adjacent→Verify_sync),DocStd(Human_ver=std_MD_best_pract;AI_ver=GenCompRules_comp_std&MUST_maintain_proper_YAML_frontmatter)|MigrationPlan:Phase1(Test_single_doc_co_located_impl),Phase2(Migrate_00_SHARED_dual_fold→co_located),Phase3(Apply_all_agent_fold_proj_wide),Phase4(Update_all_ref_doc_new_struct)|APISpec:get_context_file(filename)→if_exists(filename+".aicomp")return_comp_content:else_return_human_content;create_compressed_version(filename)→read_human_ver→apply_GenCompRules→write_filename+".aicomp";sync_versions(filename)→compare_human&comp_timestamps→upd_comp_if_human_newer|QualAssur:Reg_verify_human&comp_ver_sync,Valid_comp_ver_retain_crit_info,Obsidian_ignore_working_proper,API_intel_select_correct_ver_ctx