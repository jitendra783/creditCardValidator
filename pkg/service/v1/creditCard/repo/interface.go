package repo

type srvSt struct{

}
func NewRepoService () RepoLayer{
	return &srvSt{}
}

type RepoLayer interface{
	CompanyDetails(string ) (interface{},error)
}