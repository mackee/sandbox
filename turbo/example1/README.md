## Mojoliciousで動かすhotwire(turbo)

## 発表スライド

[https://speakerdeck.com/mackee/mojoliciousdeshu-iteli-jie-suruhotwirefalseshi-zu-mi](https://speakerdeck.com/mackee/mojoliciousdeshu-iteli-jie-suruhotwirefalseshi-zu-mi)

## 動かし方

動作確認環境: Perl 5.30.0(5.32でも動くと思われる)

[2020年にシュッとPerlの環境構築をする](https://zenn.dev/anatofuz/articles/2742225639f9f8d7bb98)などを参照し、Perl環境を整備する。system perlでの動作は確認していない。

`cpm`と`carton`っを導入する。`plenv install-cpanm`でcpanmは導入済みの前提。

```sh
$ cpnam App::cpm Carton
```

依存モジュールをインストールする。

```sh
$ cpm install
```

サーバーを起動する。

```sh
$ carton exec perl app.pl daemon
```

[http://127.0.0.1:3000](http://127.0.0.1:3000)をブラウザで開くと、デモアプリが閲覧できる。

## 注意点

### データの永続化はしていない

あくまでturboの検証のために作ったものであり、トグル状態やチャットの内容などはプロセス内にのみ保持しており、データは永続化していない。プロセスを落としてまた立ち上げた場合はデータが消えているので注意。

### XSSなどの脆弱性の検討はしていない

Mojoliciousのテンプレートレンダラーで防がれている可能性が高いが、XSSなどの脆弱性への対処は考慮していない。CSRFも現代に置いてはCORSで防御される可能性が高いが、考慮せずに作っているのでそのまま流用しないように。

## 参考にしたドキュメント等

* [Turbo Handbook](https://turbo.hotwire.dev/handbook/introduction)
* [@hotwired/turbo](https://github.com/hotwired/turbo)
* [turbo-rails gem](https://github.com/hotwired/turbo-rails)
* [remast/go\_websocket\_turbo](https://github.com/remast/go_websocket_turbo)
