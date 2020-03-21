package main

import "github.com/endw0901/go_template" // http, httpsは不要
/*
▼moduleについて
・GOPATHに縛られず、どこからでもpackageをimportできる
・基本的には、github等のレポジトリを使う

▼go build
go get => go path directoryにimportされる

▼moduleの作り方
(1)githubアカウントにrepository作成：go_math

(2)ローカルにgo_math directory作成(mkdir go_math) => 配下にcalcとgeometryフォルダを作成

(3)同フォルダにgo.mod作成
$ go mod init github.com/endw0901/go_math
go: creating new go.mod: module github.com/endw0901/go_math

(4)git init
git initで.gitの隠しファイルを作成
Initialized empty Git repository in C:/Users/masaya/go/src/go_math/.git/
※この時点でリモート接続はない
 => ls -a で隠しファイルも含めて表示可能

(5)リモートレポジトリ作成し、originとネーミング
git remote add origin https://github.com/endw0901/go_math.git
※https必要、.git必要

(6)作成したリモートリポジトリの確認
git remote -v
origin  https://github.com/endw0901/go_math.git (fetch)
origin  https://github.com/endw0901/go_math.git (push)

(7)現在のディレクトリのファイルをすべてステージに上げる
git add -A

(8)コンフィグ変数を設定し、コミットするユーザーを設定
git config user.name "endw0901"
git config user.email "endw0901@gmail.com"

(9)コミット
git commit -m "first commit"

(10)push
git push -u -f origin master

(11)バージョニング
Semantic Versioning = ユニバーサルなバージョニング仕様
x.y.z
x:major
y:minor
z:patch version number
v1.2.3
major1,minor2,patch3

git tag -a v1.0.0 -m "initial release"
git push origin master --tags

=> github上にreleaseが追加される

(12)myapp1のディレクトリ作成=>go.mod作成
1個上に戻る
cd ..
mkdir myapp1
 => main.goを作成

cd ./myapp1
ls
 => main.goが表示される

go mod init myapp1
cat go mod

(13)myapp1をrun => go.modの内容が変更されることを確認
go run main.go

$ cat go.mod
module myapp1

go 1.14

require github.com/endw0901/go_math v1.0.0

※go run コマンドは、go getコマンドを裏で実行し、imported packagesをダウンロードする
*/

func main() {
	// 1.githubに格納したpackageをimportしてcallする
	go_template.Salut()
	go_template.SayHello()
}
