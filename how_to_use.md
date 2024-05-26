# git-to-smb 使い方メモ

**このプログラムは Linux で動かすことを前提に開発したので、Windows など では動かない可能性があります**

### ステップ 1

---

git-to-smb 用のフォルダーを作り、中にバイナリを入れる

### ステップ 2

---

config.yaml を作成し、設定を入れる

**config.yaml に設定するもの一覧**
|token|username|mountedpath|
|:-:|:-:|:-:|
|GitHub のアクセストークン|保存したいユーザーの名前|samba のマウント先|

以下のように入力してください

```yaml
token: "<GitHubのアクセストークン>"
username: "<保存したいユーザーの名前>"
mountedpath: "<sambaのマウント先>"
```

### ステップ 3

---

samba を cifs でマウントする

コマンド例:

`~# mount.cifs -o sec=ntlmssp,user=<ユーザー名>,password=<パスワード>,vers=1.0,uid=1000,gid=1000 //<IP>/<保存先のフォルダ> <マウント先>`

### ステップ 4

---

git-to-smb のバイナリを実行する

history.txt, log.json, commits フォルダーなどが自動で作成されます

ログは JSON 形式で log.json に入ります

history.txt はダウンロード済みの ZIP ファイルを管理します

commits フォルダーは実際にダウンロードされた ZIP ファイルの保存先です

---

2024/05/26
