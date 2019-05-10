# Oppgave 3

Det ser ut til at mange leverandører av skytjenester leverer løsninger for talegjenkjenning eller "tale-til-tekst". Basert på dette vil vi gjøre en utforskning av noen av de største leverandørene og deres ulike implementasjoner av denne typen tjeneste. I tillegg vil vi vurdere hvor enkelt det er å integrere mot dem for oss som konsumenter og forsøke å belyse fordeler og ulemper med samtlige. Avslutningsvis vil vi oppsummere det vi har lært, og trekke frem tjenesten vi mener er best, såfremt det er mulig.

## Platformer

I arbeidet med denne oppgaven har vi forsøkt å bruke fire forskjellige "speech-to-text"-platformer.

- Google Speech API
- Google Cloud Speech-to-Text API
- Wit
- IBM Cloud Speech to Text

Først vil vi beskrive arbeidet vårt med hver av dem, deretter beskrive hvordan appen fungerer.

### Google Speech API

Scriptet i [amsehili's repository](https://github.com/amsehili/gspeech-rec) kaller et eksternt Google API med en flac-fil for å få tilbake en tekst basert på denne.

Dette APIet tilhører Chromium-teamet. Man skal [ikke bruke det til annet](https://www.chromium.org/developers/how-tos/api-keys) enn utvikling av Chromium. 

>This page is about building Chromium. If you have arrived on this page to try to get API keys for some other purpose, you should not follow the instructions on this page.

Vi gjorde likevel et forsøk å bruke det basert på API-nøkkelen som var sjekket inn i amsehili's kode, men dokumentasjonen har blitt [gjort privat](https://aminesehili.wordpress.com/2015/02/08/on-the-use-of-googles-speech-recognition-api-version-2/) og vi får en 400-feilmelding (Bad Request) ved kjøring. Etter alt dette valgte vi derfor å gå vekk fra dette APIet.

### Google Cloud Speech-to-Text API

Først, kjør:

`go get -u cloud.google.com/go/speech/apiv1`

I motsetning til Chromium-teamets API er dette åpent for bruk i andre sammenhenger. Her har vi tatt utgangspunkt i [Google's quickstart-eksempel](https://github.com/GoogleCloudPlatform/golang-samples/blob/master/speech/speech_quickstart/main.go) i Go. Vi har i tillegg lagt inn referanse til en service-legitimasjonsfil i go-koden som trengs for å autentisere seg mot Google Cloud. 

`os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "google-credentials.creds.json")`

Denne er ikke sjekket inn i repositoriet, da den er personlig og hemmelig. For å kjøre koden må man generere sin egen legitimasjonsfil i [Google's Cloud Developer Dashboard](https://console.cloud.google.com/apis/credentials), navngi den `google-credentials.creds.json` og legge den i `/ICA04/Oppgave3/google-cloud`. Prosessen er også beskrevet [her](https://cloud.google.com/speech-to-text/docs/quickstart-client-libraries#before-you-begin).

#### Autentisering mot Google Cloud

For å kompilere og kjøre denne koden forutsetter det at man har en Service account fra Google Cloud. Denne må lagres i en json-fil i `ICA04/Oppgave3/src` med navnet `google-credentials.creds.json`. Mer info om service account keys kan finnes [her](https://cloud.google.com/iam/docs/creating-managing-service-account-keys).

#### Vurdering av Google Cloud

Googles Cloud Platform Dashboard var ganske forvirrende for oss. Måten å opprette prosjekter og håndtere autentiseringsnøkler/informasjon var meget kryptisk for oss. Vi fant derimot frem til slutt. Når vi hadde dette på plass var det ganske enkelt å komme opp og kjøre. Google har som nevnt egen SDK og eksempler for Speech-to-text APIet skrevet i Go, så mye av jobben med å komme opp og kjøre var gjort for oss. De er derimot ganske pirkete på lydfilformatet (må være .flac mono), samt at man må spesifisere lydfilens egenskaper i requestene (bitrate osv.). På en annen side er Google den av leverandørene som har flest muligheter for konfigurasjon av kallene til APIet. De har for eksempel et flagg man kan sette for å skru på "profanity filter" eller "bannefilter". Noe som fungerer utmerket og sensurerer ord som klassifiseres som banneord.

### Wit

Først, kjør:

`go get -u github.com/wit-ai/wit-go`

I arbeidet med Wit tok vi utgangspunkt i koden i quickstart-eksempelet for Go som finnes [her](https://github.com/wit-ai/wit-go). Deretter opprettet vi en brukerkonto og la API-nøkkelen i en fil som vi la til i .gitignore for å unngå at den ble sjekket inn i repositoriet. Vi henter den herfra ved runtime, så hver person som skal kjøre koden må ha sin egen fil med API-nøkkel.

I de første forsøkene våre mot dette APIet brukte vi den samme lydfilen som mot Google Cloud (.flac). Dette viste seg å være problematisk, da [Wit kun støtter wav, mp3, ulaw og raw](https://wit.ai/docs/http/20160526). Vi konverterte filen til .wav for å bruke denne mot Wit-APIet. Samtidig endret vi Content-Type-headeren i kodeeksempelet til `audio/wav`.

#### Autentisering mot Wit

For å kompilere og kjøre denne koden forutsetter det at man har en Server Access Token fra Wit. Denne må lagres i en tekstfil i `ICA04/Oppgave3/src` med navnet `wit-credentials.creds.txt`.

#### Vurdering av Wit

Wit var den enkleste tjenesten å komme i gang med. Her var det bare å opprette en bruker (eller logge inn med GitHub) og så hadde man øyeblikkelig tilgangsnøkkelen man trengte for å implementere en tjenesteintegrasjon mot APIet. De hadde i tillegg en SDK og kodeeksempler i Go, noe som gjorde implementasjonen svært enkel for oss. Det var en liten ulempe at ikke .flac-filer var støttet, og dokumentasjonen for `Content-Type` headeren var noe kryptisk. Derimot helt klart den enkleste tjenesten å komme i gang med å konsumere.

### IBM Watson Speech to Text API

Etter å ha laget konto hos IBM så måtte vi opprette en Speech-to-text ressurs i deres cloud-miljø. IBM er den første leverandøren som ikke har offisielle eksempler skrevet i Go, så vi gjorde et kjapt Google-søk og fant [dette repositoriet](https://github.com/liviosoares/go-watson-sdk) som vi valgte å forsøke å ta utgangspunkt i. Det viste seg derimot at det var enklere å bare bruke http til å kalle REST-APIet direket, fremfor å ta i bruk en SDK. Dette fordi IBM bruker en enklere form for autentisering enn de øvrige.

Vi brukte det innebygde http-biblioteket i Go, med referanse i [curl-eksempelet fra IBM](https://cloud.ibm.com/docs/services/speech-to-text?topic=speech-to-text-gettingStarted#transcribe).

#### Autentisering mot IBM

For å kompilere og kjøre denne koden forutsetter det at man har en API-nøkkel hos IBM. Denne må lagres i en tekstfil i `ICA04/Oppgave3/src` med navnet `watson-credentials.creds.txt`.

#### Vurdering av IBM Watson

Det første som møtte oss etter å ha opprettet en bruker hos IBM var et dashboard som føltes veldig "corporate". Det var mange begreper som vi ikke helt forstod, så vi famlet oss frem til hvordan vi skulle skaffe oss tilgangsnøkler til APIet. I tillegg var det skuffende for oss at ikke IBM tilbyr kodeeksempler i Go. Det ser heller ikke ut til at de tilbyr kodeeksempler i noen andre konkrete språk. De har derimot et godt dokumentert REST-API som det går an å programmere mot på en enkel måte, selv med serverside-autentisering (kun Basic authentication). Vi ble derimot tvunget til å skrive opp våre egne structs for å deserialisere responsen fra APIet, noe som var litt tungvint. Vi hadde likt å se en SDK for Go.

## Vår apps funksjonalitet

Dersom alle kravene til autentisering er oppfylt kan man kjøre applikasjonen.

Applikasjonen kjøres ved å stå i `ICA04/Oppgave3/src` og kjøre `go run speech-to-text.go`. Dette vil sende `this-is-a-test.flac` som ligger i `ICA04/Oppgave3/src/sound-files` til hver av de tre APIene som vi har integrert mot og vise resultatet fra hver av dem. Som nevnt ovenfor krever samtlige autentisering, noe som ikke er del av repositoriet da dette i noen tilfeller kan påføre oss økonomiske kostnader.

## Konklusjon

De aller fleste seriøse skytjenesteleverandører har gode løsninger for å tolke tale. Alle tolket testfilene våre helt korrekt. Vi har derimot kun gjort tester på engelsk, til tross for at det også finnes støtte for andre språk. Ut fra disse vurderingene ser det ikke ut til å være noe stort som skiller de ulike leverandørene. Alle tillater gratis bruk, opp til et visst punkt. Wit ser ut til å være mer fokusert på AI og tolkning av tale mot såkalte "intents". IBM og Google ser også ut til å ha støtte for liknende tjenester, men de fremstår som bitjenester i motsetning til hos Wit, hvor dette ser ut til å være hovedfokus. Wit var helt klart den letteste å komme i gang med, mens Google hadde flest konfigurasjonsmuligheter. IBM var mest tungvint for vårt utgangspunkt da de mangler SDK og kodeeksempler i Go. Derimot var det ikke noen stor sak å implementere det vi trengte som et http-request. Dermed slapp vi en ekstra avhengighet i prosjektet vårt, noe som vi ser på som positivt.

For oss fremstår det som at det er små fordeler og ulemper med alle leverandørene vi har sett på, så vi mener ikke det er riktig å peke på en av dem som en klar "vinner". Vi holder derimot en liten knapp på Wit, ettersom denne var den enkleste å starte opp med.
