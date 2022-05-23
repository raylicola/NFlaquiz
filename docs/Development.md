## ブランチ (Branch)
* develop, master 上では作業しない (Don't work on "master" and "develop" branch.)
* develop からブランチを切ってここで作業する (Cut a branch from "develop" and work here.)
* Pull requestごとにブランチを切る ("One branch" and "one pull request" correspond one-to-one)

## フォルダ構成 (Folder Structure)
```
NFlaquiz/
　├ docker/ -> Information for Dockerfile
　├ api/    -> Backend
　├ vue     -> Frontend
　├ .env    -> environment variable
　└ docker-compose.yml
 ```
