# Git / Github
## 概要
今回ば今日プログラム開発で一般的に用いられているGit, GitHubが何なのか、どうやって使うのかを学んでいきます。

## 準備

1. [GitHubのサインアップページ](https://github.com/join)にアクセスして、アカウントを作成しましょう。
2. [設定ページ]に行き、二段階認証の設定をしてください。
3. もし第一回でGitを入れていない場合は
[ここ](../1st/readme.md/#gitの導入)
を参考に導入してください。
4. Gitの改行コードの自動置換設定を修正します。

```sh
$ git config --global core.autoCRLF false
```

4. SSHキーを登録します。

Terminal / GitBashを開き、次のコマンドを入力します。

```sh
$ mkdir .ssh
$ cd .ssh
$ ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/c/Users/USER_NAME/.ssh/id_rsa): github
Enter passphrase (empty for no passphrase): エンター
Enter same passphrase again: エンター
```

次に、次のコマンドを打ちconfigファイルを編集します。

**mac**

```sh
$ echo >> config
$ open config
```
で出てきたエディタで

```
Host github github.com
IdentityFile ~/.ssh/github
HostName github.com
User git
```

と追記し、保存します。

**Windows**

```sh
$ echo >> config
$ explorer config
```

VSCode(なければメモ帳)を選択し、そこに

```
Host github github.com
IdentityFile ~/.ssh/github
HostName github.com
User git
```

と追記し、保存します。


そうしたら、Terminal / GitBashに戻り、

```sh
$ cat github.pub
```

として出てきた`ssh-rsa`から始まる文字列をコピーして
[GitHubの登録ページ](https://github.com/settings/ssh/new)
に登録します。

最後に

```sh
$ ssh git
```
を実行したとき、

```
Hi YOUR_NAME! You've successfully authenticated, but GitHub does not provide shell access.
Connection to github.com closed.
```

と出てこれば成功です。

以上で準備は完了です。

## Gitについて

[学んで行きましょう！](./git.md)