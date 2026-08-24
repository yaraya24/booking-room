# Backend de Reservation de Salles de Reunion

### Prerequis :
- Golang 
- Sqlite3 (gcc est necessaire et il faut definir la variable d'environnement CGO_ENABLED=1)

### Installation


1. Clonez le depot
````clone git@github.com:yaraya24/booking-room.git```rr

2. Configurez la base de donnees 

```
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```
3. vous pouvez ensuite lancer le serveur avec 
```
go run cmd/main.go 
```

4. Ou vous pouvez compiler l'application en un executable qui pourra ensuite etre lance :
```
go build ./cmd
```

### Utilisation

Vous devrez utiliser l'authentification Basic Auth pour acceder a l'API.
Les utilisateurs sont listes dans le fichier setub-db.sql, ou chaque utilisateur a le mot de passe `password`.
```
Jane:password
John:password
Sarah:password
```

La creation d'une reservation se fait via l'endpoint `POST localhost:8080/bookings`
Il faut un corps de requete avec `room` et `date`. La date doit etre au format `YYYY/MM/DD`

exemple :
```
{
	"date": "2024-10-10",
	"room": "D"
}
```

La consultation des salles disponibles se fait via `GET localhost:8080/bookings?{date}` ou date est au format `YYYY/MM/DD`.

## Bugs/Problemes
1. Il y a un bug assez serieux : les utilisateurs peuvent reserver des salles qui n'existent pas. Cela vient du fait que l'application ne verifie pas l'existence de la salle avant d'effectuer la reservation et fait aveuglement confiance au client (je m'en suis rendu compte un peu trop tard).

2. Il me manque de la validation pour les requetes POST des utilisateurs

3. J'avais decide d'utiliser un fichier sql pour configurer la base de donnees, et par consequent je n'ai pas pu hacher les mots de passe. Ils sont desormais stockes en clair, ce qui n'est pas correct.

4. Nous n'avons aucune colonne meta dans nos bases de donnees, comme les horodatages de creation et de mise a jour

5. Nous ne faisons pas de ping vers la base de donnees pour verifier qu'elle fonctionne bien

6. J'aurais voulu avoir un middleware fournissant la journalisation des requetes, le request-id, le statut de la reponse, etc., mais je n'ai pas eu le temps.

7. Je n'ai pas de veritables tests d'integration : les seuls que j'ai ajoutes se trouvent dans la couche repository. La encore, faute de temps, meme si idealement cela devrait etre fait via quelque chose comme Jenkins en tant que smoke test.

## Ameliorations
1. Il aurait ete interessant d'ajouter un cache comme Redis ou un cache en memoire pour ameliorer la scalabilite de l'application. Lors de la verification des salles disponibles pour une date, on pourrait utiliser quelque chose comme un cache LRU, et lorsqu'une reservation a lieu, cette date pourrait etre mise a jour. Cela est aggrave par le fait que sqlite3 ne permet pas l'acces concurrent.

2. Ameliorer la base de donnees, en passant a mySQL ou Postgres. J'aurais pu configurer certaines options pour ameliorer les performances de sqlite, mais je n'ai pas eu le temps d'approfondir. En fin de compte, une base de donnees prete pour la production serait preferable.

3. Pour la securite/fiabilite - un limiteur de debit (rate limiter) serait egalement utile pour garantir que notre service est protege contre une utilisation intensive ou meme malveillante.
