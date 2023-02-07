// Code generated by "stringer"; DO NOT EDIT.

package privilege

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ALL-1]
	_ = x[CREATE-2]
	_ = x[DROP-3]
	_ = x[DEPRECATEDGRANT-4]
	_ = x[SELECT-5]
	_ = x[INSERT-6]
	_ = x[DELETE-7]
	_ = x[UPDATE-8]
	_ = x[USAGE-9]
	_ = x[ZONECONFIG-10]
	_ = x[CONNECT-11]
	_ = x[RULE-12]
	_ = x[MODIFYCLUSTERSETTING-13]
	_ = x[EXTERNALCONNECTION-14]
	_ = x[VIEWACTIVITY-15]
	_ = x[VIEWACTIVITYREDACTED-16]
	_ = x[VIEWCLUSTERSETTING-17]
	_ = x[CANCELQUERY-18]
	_ = x[NOSQLLOGIN-19]
	_ = x[EXECUTE-20]
	_ = x[VIEWCLUSTERMETADATA-21]
	_ = x[VIEWDEBUG-22]
	_ = x[BACKUP-23]
	_ = x[RESTORE-24]
	_ = x[EXTERNALIOIMPLICITACCESS-25]
	_ = x[CHANGEFEED-26]
}

const _Kind_name = "ALLCREATEDROPGRANTSELECTINSERTDELETEUPDATEUSAGEZONECONFIGCONNECTRULEMODIFYCLUSTERSETTINGEXTERNALCONNECTIONVIEWACTIVITYVIEWACTIVITYREDACTEDVIEWCLUSTERSETTINGCANCELQUERYNOSQLLOGINEXECUTEVIEWCLUSTERMETADATAVIEWDEBUGBACKUPRESTOREEXTERNALIOIMPLICITACCESSCHANGEFEED"

var _Kind_index = [...]uint16{0, 3, 9, 13, 18, 24, 30, 36, 42, 47, 57, 64, 68, 88, 106, 118, 138, 156, 167, 177, 184, 203, 212, 218, 225, 249, 259}

func (i Kind) String() string {
	i -= 1
	if i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
