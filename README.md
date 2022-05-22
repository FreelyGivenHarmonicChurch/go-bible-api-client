# go-bible-api-client

asdf plugin add golang
asdf install

go mod init freelygivenharmonic.church/go-bible-api-client/v2

go build .
./go-bible-api-client

export $(grep -v '^#' .env | xargs)


King James
id: de4e12af7f28f599-02

curl -X GET "https://api.scripture.api.bible/v1/bibles/de4e12af7f28f599-02" -H  "accept: application/json" -H  "api-key: ******"

{
  "data": {
    "id": "de4e12af7f28f599-02",
    "dblId": "de4e12af7f28f599",
    "relatedDbl": null,
    "name": "King James (Authorised) Version",
    "nameLocal": "King James Version",
    "abbreviation": "engKJV",
    "abbreviationLocal": "KJV",
    "description": "Protestant",
    "descriptionLocal": "Protestant",
    "language": {
      "id": "eng",
      "name": "English",
      "nameLocal": "English",
      "script": "Latin",
      "scriptDirection": "LTR"
    },
    "countries": [
      {
        "id": "GB",
        "name": "United Kingdom",
        "nameLocal": "United Kingdom"
      }
    ],
    "type": "text",
    "updatedAt": "2022-01-07T15:05:57.000Z",
    "copyright": "PUBLIC DOMAIN except in the United Kingdom, where a Crown Copyright applies to printing the KJV. See http://www.cambridge.org/about-us/who-we-are/queens-printers-patent",
    "info": "<p>This historical translation of the Holy Bible is brought to you in quality digital format by the <a href=\"https://crosswire.org/\">Crosswire Bible Society</a> and <a href=\"https://eBible.org\">eBible.org</a>.</p>",
    "audioBibles": []
  }
}