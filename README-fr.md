# acli
Une interface en ligne de commande pour interagir avec les modèles IA d'OpenAI.

## Prérequis

### Obtenir une clé API OpenAI

Inscrivez-vous sur le site web d'OpenAI API : https://openai.com/api/. Après vous être connecté, créez une clé API à cette adresse URL : https://beta.openai.com/account/api-keys.

## Utilisation

Utilisez acli pour les discussions ou la résolution de tâches complexes. Exemples :
* `acli "GPT peut-il m'aider pour les tâches quotidiennes en ligne de commande ?"`
* `acli "[Description complexe de la demande de fonctionnalité pour bash/javascript/python/etc...]"`

Utilisez la fonction howto pour obtenir rapidement des réponses en une ligne et le mode interactif. Exemples :
* `howto openssl tester l'expiration SSL de github.com`
* `howto "trouver tous les fichiers de plus de 30 Mo"`

Utilisez fix pour corriger rapidement les fautes de frappe. Exemples :
* [Exécutez une commande avec une faute de frappe comme 'rrm', 'lls', 'cd..', etc.]
* Ensuite, tapez `fix` et obtenez la commande corrigée prête à être exécutée

## Installer

### Installer avec un script

Exécutez :
```
curl -sSLO https://raw.githubusercontent.com/jseguillon/acli/main/get.sh && \
bash get.sh
```

### Ou installer manuellement
Accédez à la [page des versions](https://github.com/jseguillon/acli/releases), trouvez le binaire approprié pour votre système. Téléchargez-le, installez-le où vous le souhaitez et utilisez chmod +x dessus. Exemple :

```
sudo curl -SL [release_url] -o /usr/local/bin/acli
sudo chmod +x /usr/local/bin/acli
```

Ajoutez la configuration dans n'importe quel fichier .rc de votre choix :

```
CHAT_GPT_API_KEY="XXXXX"

alias fix='eval $(acli --script fixCmd "$(fc -nl -1)" $?)'
howto() { h="$@"; eval $(acli --script howCmd "$h") ; }
```

## License

This program is licensed under the MIT License. See [LICENSE](LICENSE) for more details.
