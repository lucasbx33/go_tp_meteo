# TP Météo

## Correspondance des données entre JSON et XML

| Donnée                   | JSON                                                                                   | XML                                                                                          |
|--------------------------|----------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------|
| **Pays**                 | "country": "France"                                                                    | ```<station country="FR">```                                                                 |
| **Coordonnées**          | "location": {"latitude": 44.8333,"longitude": -0.7}                                    | ```<station><coordinates lat="44.8333" lon="-0.7"/>```                                       |
| **Altitude**             | "altitude_m": 47                                                                       | ```<station><coordinates altitude="47"/>```                                                  |
| **Modèle de capteur**    | "device": {"type": "AWS-3000","manufacturer": "Vaisala","installed_on": "2015-12-09"}, | ```<station><hardware vendor="Vaisala" model="AWS-3000" since="2015-12-09"/>```              |
| **Température**          | "observations": [{"temperature_celsius": 3.3}]                                         | ```<station><observations><observation><measure type="temperature" unit="C">3.3</measure>``` |
| **Conditions ciel**      | "observations": [{"conditions": "clear"}]                                              | ```<station><observations><observation sky="clear">```                                       |
| **Vent**                 | "observations": [{"wind": {"speed_kmh": 48.1,"direction_deg": 279}}]                   | ```<station><observations><observation><wind speed="48.1" direction="279"/>```               |
| **Notes (optionnelles)** | "observations": [{"notes": null}]                                                      | ```<station><observations><observation><note>Calibrage capteur effectué la veille</note>```  |
