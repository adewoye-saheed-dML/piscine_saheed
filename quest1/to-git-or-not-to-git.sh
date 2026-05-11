curl https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[]|select(.id==170).name,select(.id==170).powerstats.power, select(.id==170).appearance.gender'

# |select(.powerstats.power)