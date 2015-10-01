# isolat
Isolat-tjenesten tar i mot en fange til isolasjon over en gitt tid. Deretter slippes fangen ut av isolat og sendes videre.

Docker oppstart
<code>
  docker run -d -p 9998:9998 --name isolat mesanfagark/reaktiv-isolat
</code>

Isolat svarer på url:
<code>
http://localhost:9998/isolat/
</code>

For å sende en fange til isolat, sendes en HTTP POST til http://localhost:9998/isolat/.
Den bryr seg foreløbig ikke om eventuelle headere...
```
{
  "FangeTilIsolat":
    {
      "Id":"1ES532KD1",
      "Navn":"Albert Åbert"
    },
  "IsoleringsTid":5,
  "CallbackUrl":"http://dummy.url/",
  "Method":"GET"
}
```

IsoleringsTid - tid i sekunder som fangen blir liggende i isolat
CallbackUrl - url fangen sendes til når isoleringstiden er ute.
Method - HTTP metode brukt.

TODO: Må vel legge til noen headere i callbacket.


