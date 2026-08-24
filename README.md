# Backend per la Prenotazione di Sale Riunioni

### Requisiti:
- Golang 
- Sqlite3 (gcc e' necessario e impostare la variabile d'ambiente CGO_ENABLED=1)

### Installazione


1. Clona il repository
````clone git@github.com:yaraya24/booking-room.git```rr

2. Configura il database 

```
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```
3. puoi quindi avviare il server usando 
```
go run cmd/main.go 
```

4. Oppure puoi compilare l'app come eseguibile che puo' poi essere avviato:
```
go build ./cmd
```

### Utilizzo

Sara' necessario usare la Basic Auth per accedere all'API.
Gli utenti sono elencati nel file setub-db.sql, dove ogni utente ha la password `password`.
```
Jane:password
John:password
Sarah:password
```

Per creare una prenotazione si usa l'endpoint `POST localhost:8080/bookings`
E' necessario un body con `room` e `date`. La data deve essere nel formato `YYYY/MM/DD`

esempio:
```
{
	"date": "2024-10-10",
	"room": "D"
}
```

Per visualizzare le sale disponibili si puo' usare `GET localhost:8080/bookings?{date}` dove date e' nel formato `YYYY/MM/DD`.

## Bug/Problemi
1. C'e' un bug piuttosto serio: gli utenti possono prenotare sale che non esistono. Questo accade perche' l'app non verifica l'esistenza della sala prima di effettuare la prenotazione e si fida ciecamente del client (me ne sono accorto un po' troppo tardi).

2. Manca della validazione per le richieste POST degli utenti

3. Avevo deciso di usare un file sql per configurare il database e di conseguenza non sono riuscito ad applicare l'hashing alle password. Ora sono memorizzate in chiaro, il che non va bene.

4. Non abbiamo colonne meta nei nostri database come i timestamp di creazione e aggiornamento

5. Non facciamo il ping al database per assicurarci che sia effettivamente in esecuzione

6. Avrei voluto avere un middleware che fornisse il logging della richiesta, il request-id, lo status della risposta, ecc., ma non ho avuto tempo.

7. Non ho scritto veri test di integrazione: gli unici che ho aggiunto sono a livello di repository. Anche in questo caso per motivi di tempo, anche se idealmente andrebbe fatto tramite qualcosa come Jenkins come smoke test.

## Miglioramenti
1. Sarebbe stato bello aggiungere una cache come Redis o una cache in memoria per migliorare la scalabilita' dell'applicazione. Quando si controllano le sale disponibili per una data, si potrebbe usare qualcosa come una cache LRU e, quando avviene una prenotazione, quella data potrebbe essere aggiornata. Questo e' aggravato dal fatto che sqlite3 non permette l'accesso concorrente.

2. Migliorare il database, passando a mySQL o Postgres. Avrei potuto impostare alcune opzioni per migliorare le prestazioni di sqlite ma non ho avuto tempo di approfondire. In definitiva, sarebbe preferibile un database pronto per la produzione.

3. Per sicurezza/affidabilita' - anche un rate limiter sarebbe utile per assicurare che il nostro servizio sia protetto da un uso intenso o addirittura malevolo.
