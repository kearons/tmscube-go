package constant

var MessageTypes = [5]string{
	"course", "grab", "edu", "other", "invk",
}

var SpecialList = map[string]map[string]int{
	"baku_class_new"				: {"type" : 79, "showType" : 1},
	"baku_class_time_change"        : {"type" : 79, "showType" : 2},
	"baku_class_cancel"             : {"type" : 79, "showType" : 3},
	"baku_confirm_class_success"    : {"type" : 89, "showType" : 4},
	"baku_confirm_class_failed"     : {"type" : 89, "showType" : 5},
	// "baku_class_start_remind"    : {"type" : 89, "showType" : 6},
	"baku_class_new_v2"             : {"type" : 79, "showType" : 1},
	"baku_class_time_change_v2"     : {"type" : 79, "showType" : 2},
	// "baku_class_cancel_v2"          : {"type" : 79, "showType" : 3},//2020-11-23 15:41:48去掉
	"baku_class_duration_change_v2" : {"type" : 79, "showType" : 2},
	"baku_teacher_orientate_invite" : {"type" : 149, "showType" : 10},//邀课
	"tms_one_absent_notice"         : {"type" : 149, "showType" : 8},
	"tms_two_absent_notice"         : {"type" : 149, "showType" : 8},
	"tms_absent_finish_deal"        : {"type" : 149, "showType" : 9},
	"tms_class_absent_in_one"       : {"type" : 149, "showType" : 8},
	"tms_class_absent_in_two"       : {"type" : 149, "showType" : 8},
	"baku_fix_class_arrange_edit"   : {"type" : 79, "showType" : 2},
	"baku_fix_class_arrange_new"    : {"type" : 79, "showType" : 1},
}

var DefaultPushType = map[string]int{
	"course": 30,
	"grab":   39,
	"edu":    31,
	"other":  32,
	"invk":   33,
}
