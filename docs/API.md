## APIの仕様

| URL     | メソッド | やること                       | 受け取る値                        | 返す値      | 
| ------- | -------- | ------------------------------ | --------------------------------- | ----------- | 
| /signup | POST     | ユーザーの新規登録             | email, password, password_confirm |             | 
| /login  | POST     | ログイン                       | email, password                   |Cookieにセットしたトークン<br>{"jwt": tokenString}| 
| /logout | GET      | ログアウト                     |                                   |             | 

### 共通事項
返り値
* msg : ちゃんとアクセスできてるかテスト用
* err_msg : エラーメッセージ
  * 返り値に "err_msg" が含まれる場合, 何らかのエラーが発生している
