package builder

func DirectBuildPerson(b PersonBuider) {
	b.BuildHead()
	b.BuildBody()
	b.BuildArms()
	b.BuildLegs()
}
