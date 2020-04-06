# [Talk Go](https://talkgo.org/)
[![Build Status](https://travis-ci.org/talk-go/night.svg?branch=master)](https://travis-ci.org/talk-go/night) [![Go Report Card](https://goreportcard.com/badge/github.com/talk-go/night)](https://goreportcard.com/report/github.com/talk-go/night) [![GitHub stars](https://img.shields.io/github/stars/talk-go/night.svg?label=Stars)](https://github.com/talk-go/night) [![GitHub forks](https://img.shields.io/github/forks/talk-go/night.svg?label=Fork)](https://github.com/talk-go/night) [![All Contributors](https://img.shields.io/badge/all_contributors-48-orange.svg?style=flat-square)](#contributors) [![Documentation](https://godoc.org/github.com/talk-go/night?status.svg)](http://godoc.org/github.com/talk-go/night) [![Coverage Status](https://coveralls.io/repos/github/talk-go/night/badge.svg?branch=master)](https://coveralls.io/github/talk-go/night?branch=master) [![GitHub issues](https://img.shields.io/github/issues/talk-go/night.svg?label=Issue)](https://github.com/talk-go/night/issues) [![license](https://img.shields.io/github/license/talk-go/night.svg)](https://github.com/talk-go/night/blob/master/LICENSE)

<img src="https://raw.githubusercontent.com/talk-go/night/master/static/images/2018-12-11-night-reading-go.jpg" width="400px;"/>

#### *Lies das in [anderen Sprachen](Translations.md).*

[🇨🇳](README.md)
[🇭🇰](README.cht.md)
[🇺🇸](README.en.md)
[🇩🇪](README.de.md)

## Sterngucker im Zeitablauf

[![Stargazers Over Time](https://starcharts.herokuapp.com/talk-go/night.svg)](https://starcharts.herokuapp.com/talk-go/night)

[reading-go Sternenhistory und Statistiken](https://seladb.github.io/StarTrack-js/#/preload?r=talk-go,night)

Go Studieren und Teilen:

- [Studien zum Go in der Nacht](https://github.com/talk-go/night/labels/Go%20%E5%A4%9C%E8%AF%BB)

*Wir werden einen Abend pro Woche ein Treffen vereinbaren, um den Go-Quellcode zu lesen. Wenn du ein Neuling bist, kannst du ihn hier besuchen. **[Die Art und Weise des Lernens Go](https://github.com/talk-go/read)**.*

>Unsere Studie umfasst: Go Standard Bibliothek, Open Source Projekt.

- [Tägliches Lernen](https://github.com/talk-go/night/labels/%E6%AF%8F%E6%97%A5%E9%98%85%E8%AF%BB)

*Wenn du an diesen Meetings teilnimmst, bedeutet das, dass du etwas ändern willst, und du hast dir versprochen, dich zu ändern. Allerdings musst du das verstehen: Wir können diskutieren, miteinander teilen, gemeinsam etwas beitragen, aber wir haben keine Verpflichtung, dich zu unterrichten oder zu führen.*

- [Tägliche Frage](https://github.com/talk-go/night/labels/%E6%AF%8F%E6%97%A5%E4%B8%80%E9%97%AE)

- [Projektpraktiken](https://github.com/talk-go/night/labels/Go%20%E9%A1%B9%E7%9B%AE%E5%AE%9E%E8%B7%B5)
- [Entwicklung im Gin](https://github.com/talk-go/night/labels/Gin%20%E5%BC%80%E5%8F%91)

----

<br>

| ![notification](/static/images/bell-outline-badged.svg)Vorankündigung |
| :----------------------------------------------------------: |
|          k8s Source Code - Scheduler 4/13 John           |
| github.com/golang/sync -> errgroup、syncmap ... source code reading |

----

## Fallbeispiel

- [x] strings
- [x] strconv
- [x] testing
- [x] net/http
- [x] sync
- [x] flag
- [x] etcd/raft
- [x] defer
- [x] context
- [ ] golang/sync
- [ ] kubernetes

## Review Netzplatz

- [Talk Go (YouTuBe)](https://youtube.com/c/talkgo_night)
- [Talk Go (Bilibili)](https://space.bilibili.com/326749661)
- [Talk Go (Tencent)](https://v.qq.com/vplus/e05f55c8ca5e36f8e370ba49c9e883e0)

## Unser Ziel

Wir hoffen, dass alle, die teilgenommen haben, "Go" tief verstehen können und schneller ein Senior Gopher werden. Wir wünschen uns auch, dass jeder Wissen und Erfahrungen sammeln kann.

## Unser Methode

Wir werden einen Führer haben, der den Go-Quellcode jedem vorstellt und die schwierigen Algorithmen und Programmiertricks studiert. Wenn es Probleme oder Zweifel gibt, können wir gemeinsam recherchieren und die Antwort finden.

>Lese-Prinzip: Wählen Sie ein Paket aus und lesen Sie eine `.go`-Datei von oben nach unten. Wir werden die Datei Schritt für Schritt überprüfen, aber wir werden nicht in eine Logik eintauchen, die außerhalb des zu prüfenden Pakets liegt.

## Unser Geist

Open Source! Open Source! Open Source! (**sondern, nochmals gesagt und dreimal gesagt** :-) )

>Wir hoffen, dass sich alle Interessierten uns anschließen und dazu beitragen, "Studien zum Go in der Nacht" zu einer Open-Source-Community zu machen, von der alle profitieren können.

## Wie kann man beitreten?

Wenn Sie der WeChat-Gruppe beitreten möchten, suchen Sie nach `night_reading_go` und geben Sie dann Ihren Namen, Ihr Unternehmen, Ihre Berufsbezeichnung und Ihre Verantwortung ein. Bitte erwähnen Sie auch, dass Sie uns auf Github gefunden haben.

Für diejenigen unter Ihnen, die Slack nutzen möchten, ist hier die Anfahrtsbeschreibung: [Talk Go Slack](https://join.slack.com/t/talkgo/shared_invite/zt-89zh1000-KX2tZ6l~FSNP14Oy2B~onQ)

## Night Reading Go SIG members

>Thanks for you effort.

- [杨文 yangwenmai](https://github.com/yangwenmai)
- [欧长坤 changkun](https://github.com/changkun)
- [FelixSeptem](https://github.com/FelixSeptem)
- [饶全成 qcrao](https://github.com/qcrao)
- [煎鱼 - EDDYCJY](https://github.com/EDDYCJY)

----

## Beitragsrichtlinien

```sh
├── reading   // Go Quellcode
├── discuss   // Zusammenfassung der täglichen Diskussion
├── articles  // Technische Artikel
├── interview // Go Interview
└── other     // Andere
```

- [Beitragsrichtlinien](https://github.com/talk-go/night/blob/master/CONTRIBUTING.md)
- *[Git Commit Konventionen](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit?pref=2&pli=1#)*

## Mitwirkende

Wir schätzen jeden, der zu diesem Projekt beiträgt, sei es bei der Erstellung von Themen oder Pull-Requests (PRs), sei es einfach nur bei der Community. Wenn Sie eine PR beitragen, die wir akzeptieren, wird Ihr Name hier aufgeführt.

So können Sie eine PR erstellen:

- Installieren `npm install -g --save-dev all-contributors-cli`
- `sh gen_contributors.sh`

Es gibt viele Arten von Beiträgen, wie "code", "ideas", "reviews", "talks" oder "tutorials". Sie können `.all-contributors` ändern.

Vielen Dank an diese wunderbaren Menschen ([emoji key](https://github.com/kentcdodds/all-contributors#emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
| [<img src="https://avatars3.githubusercontent.com/u/1710912?v=4" width="100px;"/><br /><sub><b>maiyang</b></sub>](https://github.com/yangwenmai)<br />[💻](https://github.com/talk-go/night/commits?author=yangwenmai "Code") [🤔](#ideas-yangwenmai "Ideas, Planning, & Feedback") [👀](#review-yangwenmai "Reviewed Pull Requests") [📢](#talk-yangwenmai "Talks") [✅](#tutorial-yangwenmai "Tutorials") | [<img src="https://avatars1.githubusercontent.com/u/16773339?v=4" width="100px;"/><br /><sub><b>Simple Min</b></sub>](https://github.com/mougeCM)<br />[💻](https://github.com/talk-go/night/commits?author=mougeCM "Code") | [<img src="https://avatars3.githubusercontent.com/u/35653599?v=4" width="100px;"/><br /><sub><b>kenny</b></sub>](https://github.com/yuhao5)<br />[💻](https://github.com/talk-go/night/commits?author=yuhao5 "Code") | [<img src="https://avatars2.githubusercontent.com/u/13843868?v=4" width="100px;"/><br /><sub><b>charnlsxy</b></sub>](https://github.com/charnlsxy)<br />[💻](https://github.com/talk-go/night/commits?author=charnlsxy "Code") | [<img src="https://avatars3.githubusercontent.com/u/11901298?v=4" width="100px;"/><br /><sub><b>AceDarkknight</b></sub>](https://github.com/AceDarkknight)<br />[💻](https://github.com/talk-go/night/commits?author=AceDarkknight "Code") | [<img src="https://avatars2.githubusercontent.com/u/3014297?v=4" width="100px;"/><br /><sub><b>Data</b></sub>](https://github.com/gnuos)<br />[💻](https://github.com/talk-go/night/commits?author=gnuos "Code") | [<img src="https://avatars0.githubusercontent.com/u/2876745?v=4" width="100px;"/><br /><sub><b>侯名</b></sub>](https://github.com/KISSMonX)<br />[💻](https://github.com/talk-go/night/commits?author=KISSMonX "Code") |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| [<img src="https://avatars0.githubusercontent.com/u/12060175?v=4" width="100px;"/><br /><sub><b>dumliu01</b></sub>](https://github.com/dumliu01)<br />[💻](https://github.com/talk-go/night/commits?author=dumliu01 "Code") | [<img src="https://avatars0.githubusercontent.com/u/1411282?v=4" width="100px;"/><br /><sub><b>hlily2005</b></sub>](https://github.com/hlily2005)<br />[💻](https://github.com/talk-go/night/commits?author=hlily2005 "Code") | [<img src="https://avatars0.githubusercontent.com/u/16982786?v=4" width="100px;"/><br /><sub><b>fenggolang</b></sub>](https://github.com/fenggolang)<br />[💻](https://github.com/talk-go/night/commits?author=fenggolang "Code") | [<img src="https://avatars3.githubusercontent.com/u/10174178?v=4" width="100px;"/><br /><sub><b>henrylee2cn</b></sub>](https://github.com/henrylee2cn)<br />[💻](https://github.com/talk-go/night/commits?author=henrylee2cn "Code") | [<img src="https://avatars0.githubusercontent.com/u/1336914?v=4" width="100px;"/><br /><sub><b>shaqsnake</b></sub>](https://github.com/shaqsnake)<br />[💻](https://github.com/talk-go/night/commits?author=shaqsnake "Code") | [<img src="https://avatars0.githubusercontent.com/u/5728787?v=4" width="100px;"/><br /><sub><b>tbwisk</b></sub>](https://github.com/TBWISK)<br />[💻](https://github.com/talk-go/night/commits?author=TBWISK "Code") | [<img src="https://avatars3.githubusercontent.com/u/416141?v=4" width="100px;"/><br /><sub><b>Huang ChuanTong</b></sub>](https://github.com/toontong)<br />[💻](https://github.com/talk-go/night/commits?author=toontong "Code") |
| [<img src="https://avatars3.githubusercontent.com/u/10513552?v=4" width="100px;"/><br /><sub><b>The notes of SQL optimize </b></sub>](https://github.com/zhongxuan123)<br />[💻](https://github.com/talk-go/night/commits?author=zhongxuan123 "Code") | [<img src="https://avatars2.githubusercontent.com/u/29008269?v=4" width="100px;"/><br /><sub><b>zhouxinxin19920802</b></sub>](https://github.com/zhouxinxin19920802)<br />[💻](https://github.com/talk-go/night/commits?author=zhouxinxin19920802 "Code") | [<img src="https://avatars2.githubusercontent.com/u/20811449?v=4" width="100px;"/><br /><sub><b>macaria</b></sub>](https://github.com/macaria)<br />[💻](https://github.com/talk-go/night/commits?author=macaria "Code") | [<img src="https://avatars3.githubusercontent.com/u/15226239?v=4" width="100px;"/><br /><sub><b>Dennis</b></sub>](http://github.com/DennisMao)<br />[💻](https://github.com/talk-go/night/commits?author=DennisMao "Code") | [<img src="https://avatars1.githubusercontent.com/u/2696746?v=4" width="100px;"/><br /><sub><b>orangleliu</b></sub>](http://blog.csdn.net/orangleliu)<br />[💻](https://github.com/talk-go/night/commits?author=orangle "Code") | [<img src="https://avatars1.githubusercontent.com/u/21693162?v=4" width="100px;"/><br /><sub><b>HarbinZhang</b></sub>](https://github.com/HarbinZhang)<br />[💻](https://github.com/talk-go/night/commits?author=HarbinZhang "Code") | [<img src="https://avatars1.githubusercontent.com/u/7344921?v=4" width="100px;"/><br /><sub><b>LiMingji</b></sub>](https://github.com/SwanSpouse)<br />[💻](https://github.com/talk-go/night/commits?author=SwanSpouse "Code") |
| [<img src="https://avatars0.githubusercontent.com/u/22164927?v=4" width="100px;"/><br /><sub><b>wintersnow</b></sub>](https://mickey0524.github.io/)<br />[💻](https://github.com/talk-go/night/commits?author=mickey0524 "Code") | [<img src="https://avatars2.githubusercontent.com/u/44076738?v=4" width="100px;"/><br /><sub><b>zhuzhenfeng</b></sub>](https://github.com/zhuzhenfeng-finogeeks)<br />[💻](https://github.com/talk-go/night/commits?author=zhuzhenfeng-finogeeks "Code") | [<img src="https://avatars2.githubusercontent.com/u/6065007?v=4" width="100px;"/><br /><sub><b>徐佳军</b></sub>](https://xujiajun.cn)<br />[💻](https://github.com/talk-go/night/commits?author=xujiajun "Code") | [<img src="https://avatars0.githubusercontent.com/u/6884499?v=4" width="100px;"/><br /><sub><b>nicho</b></sub>](https://github.com/NichoZhang)<br />[💻](https://github.com/talk-go/night/commits?author=NichoZhang "Code") | [<img src="https://avatars1.githubusercontent.com/u/17244565?v=4" width="100px;"/><br /><sub><b>Weifeng Wang</b></sub>](https://www.btxiaowei.net)<br />[💻](https://github.com/talk-go/night/commits?author=qclaogui "Code") | [<img src="https://avatars3.githubusercontent.com/u/6748475?v=4" width="100px;"/><br /><sub><b>John Deng</b></sub>](https://hiboot.hidevops.io)<br />[💻](https://github.com/talk-go/night/commits?author=john-deng "Code") | [<img src="https://avatars0.githubusercontent.com/u/17971291?v=4" width="100px;"/><br /><sub><b>赵吉彤</b></sub>](http://jeasonstudio.github.io)<br />[💻](https://github.com/talk-go/night/commits?author=jeasonstudio "Code") |
| [<img src="https://avatars0.githubusercontent.com/u/3946563?v=4" width="100px;"/><br /><sub><b>YING ZOU</b></sub>](http://zouying.is)<br />[💻](https://github.com/talk-go/night/commits?author=xpzouying "Code") | [<img src="https://avatars3.githubusercontent.com/u/6278792?v=4" width="100px;"/><br /><sub><b>zsy619</b></sub>](https://github.com/zsy619)<br />[💻](https://github.com/talk-go/night/commits?author=zsy619 "Code") | [<img src="https://avatars2.githubusercontent.com/u/11313960?v=4" width="100px;"/><br /><sub><b>杨帆</b></sub>](https://blog.yangfan21.cn/)<br />[💻](https://github.com/talk-go/night/commits?author=yangfan21 "Code") | [<img src="https://avatars3.githubusercontent.com/u/9268902?v=4" width="100px;"/><br /><sub><b>HundredLee</b></sub>](https://blog.sodroid.com)<br />[💻](https://github.com/talk-go/night/commits?author=hundredlee "Code") | [<img src="https://avatars3.githubusercontent.com/u/1733903?v=4" width="100px;"/><br /><sub><b>mlboy</b></sub>](https://github.com/mlboy)<br />[💻](https://github.com/talk-go/night/commits?author=mlboy "Code") | [<img src="https://avatars1.githubusercontent.com/u/28711504?v=4" width="100px;"/><br /><sub><b>fish</b></sub>](https://liangyuanpeng.netlify.com/)<br />[💻](https://github.com/talk-go/night/commits?author=liangyuanpeng "Code") | [<img src="https://avatars2.githubusercontent.com/u/20315934?v=4" width="100px;"/><br /><sub><b>时小光</b></sub>](https://laji.cx)<br />[💻](https://github.com/talk-go/night/commits?author=abcdsxg "Code") |
| [<img src="https://avatars1.githubusercontent.com/u/44582639?v=4" width="100px;"/><br /><sub><b>Ziyi Yan</b></sub>](https://ziyi-yan.github.io)<br />[💻](https://github.com/talk-go/night/commits?author=ziyi-yan "Code") | [<img src="https://avatars2.githubusercontent.com/u/11867562?s=460&v=4" width="100px;"/><br /><sub><b>李朋飞</b></sub>](https://github.com/lpflpf)<br />[💻](https://github.com/talk-go/night/commits?author=lpflpf "Code") | [<img src="https://avatars0.githubusercontent.com/u/13746731?v=4" width="100px;"/><br /><sub><b>煎鱼</b></sub>](https://github.com/EDDYCJY)<br />[💻](https://github.com/talk-go/night/commits?author=EDDYCJY "Code") | [<img src="https://avatars2.githubusercontent.com/u/3349949?v=4" width="100px;"/><br /><sub><b>Wang Fei</b></sub>](https://github.com/feirie)<br />[💻](https://github.com/talk-go/night/commits?author=feirie "Code") | [<img src="https://avatars0.githubusercontent.com/u/8568271?v=4" width="100px;"/><br /><sub><b>742161455</b></sub>](https://github.com/jukylin)<br />[💻](https://github.com/talk-go/night/commits?author=jukylin "Code") | [<img src="https://avatars1.githubusercontent.com/u/3310967?s=460&v=4" width="100px;"/><br /><sub><b>feifeiiiiiiiiiii</b></sub>](https://github.com/feifeiiiiiiiiiii)<br />[💻](https://github.com/talk-go/night/commits?author=feifeiiiiiiiiiii "Code") | [<img src="https://avatars0.githubusercontent.com/u/15921519?v=4" width="100px;"/><br /><sub><b>崔爽</b></sub>](https://github.com/cuishuang)<br />[💻](https://github.com/talk-go/night/commits?author=cuishuang "Code") |
| [<img src="https://avatars1.githubusercontent.com/u/16793420?v=4" width="100px;"/><br /><sub><b>jasonxie</b></sub>](http://www.techclone.cn)<br />[💻](https://github.com/talk-go/night/commits?author=JasonRD "Code") | [<img src="https://avatars0.githubusercontent.com/u/9942270?v=4" width="100px;"/><br /><sub><b>haoc7</b></sub>](http://cuihao.fun)<br />[💻](https://github.com/talk-go/night/commits?author=itcuihao "Code") | [<img src="https://avatars3.githubusercontent.com/u/24445731?s=400&v=4" width="100px;"/><br /><sub><b>鱼乐</b></sub>](https://blog.wangriyu.wang/)<br />[💻](https://github.com/talk-go/night/commits?author=wangriyu "Code") | [<img src="https://avatars2.githubusercontent.com/u/16516151?s=400&v=4" width="100px;"/><br /><sub><b>Littlesqx</b></sub>](https://littlesqx.github.io/)<br />[💻](https://github.com/talk-go/night/commits?author=Littlesqx "Code") | [<img src="https://avatars1.githubusercontent.com/u/31753706?s=400&v=4" width="100px;"/><br /><sub><b>mchangxin</b></sub>](https://github.com/mchangxin)<br />[💻](https://github.com/talk-go/night/commits?author=mchangxin "Code") | [<img src="https://avatars0.githubusercontent.com/u/32830059?s=400&v=4" width="100px;"/><br /><sub><b>Hokkaitao</b></sub>](https://github.com/Hokkaitao)<br />[💻](https://github.com/talk-go/night/commits?author=Hokkaitao "Code") |
<!-- ALL-CONTRIBUTORS-LIST:END -->

Dieses Projekt folgt der Spezifikation [all-contributors](https://github.com/kentcdodds/all-contributors). Alle Beiträge sind willkommen!
