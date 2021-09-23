package dao

import "github.com/google/wire"

// RepoSet repo injection
var DaoSet = wire.NewSet(
	UserSet,
) // end
