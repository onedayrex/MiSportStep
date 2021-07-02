module "MiSportStep"

require (
	sport v0.0.0
)

replace (
	sport v0.0.0 => ./src
)