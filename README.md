# Analytics MX — API

Backend de **Analytics MX**, una plataforma de encuestas hecha por encargo. Expone la API que usan el panel web y la app móvil de captura en campo.

## Qué resuelve

Una encuesta no se programa: se define. El servicio guarda formularios armados con preguntas reutilizables, de modo que la misma pregunta sirve en varias encuestas, y cada formulario decide cuáles son obligatorias y cuáles sólo aparecen si otra se respondió de cierta forma. Las respuestas se guardan con la ubicación desde donde se levantaron, que es lo que permite leer un resultado por zona y no sólo en total.

## Modelo

| Tabla | Para qué |
|---|---|
| `forms` | La encuesta: nombre y descripción |
| `questions` | El catálogo de preguntas, independiente de la encuesta |
| `forms_questions` | Qué preguntas lleva cada encuesta, si son obligatorias y de cuál dependen |
| `answers` | La respuesta, ligada al formulario y a la pregunta |
| `locations` | Latitud y longitud de donde se levantó |

## Cómo funciona

Las consultas no se escriben a mano: se declaran en SQL bajo `db/queries` y **sqlc** genera el código Go con sus tipos, así que un cambio en el esquema rompe la compilación en lugar de fallar en tiempo de ejecución.

Las migraciones van embebidas en el binario con `embed.FS` y se aplican solas al arrancar, de modo que desplegar y migrar son el mismo paso.

## Stack

Go · chi · pgx · PostgreSQL · sqlc · dbmate

## Correrlo

```bash
export DATABASE_URL=postgres://usuario:clave@host:5432/base
go run .
```
