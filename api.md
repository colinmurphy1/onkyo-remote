# Onkyo Controller API Reference

## Power

**Check power status**  
`GET /api/power`

Will return `status` set to `true` if on, `false` if off.

    200 OK
    {
        "data": {
            "status": true
        },
        "message": "OK"
    }

**Turn receiver on or off**  
ON: `GET /api/power/set/on`, OFF: `GET /api/power/set/off`

Will return `true` if the receiver is now on, `false` if off.

    200 OK
    {
        "data": {
            "status": true
        },
        "message": "OK"
    }

---

## Volume

### Get volume level

`GET /api/volume`

*Not yet implemented*

### Set volume level

`PUT /api/volume/set/10`

    200 OK
    {
        "data": {
            "level": 10
        },
        "message": "OK"
    }


### Mute

*Not yet implemented* 

---

## Tuning

*Not yet implemented* 


Get frequency: `GET /api/tune`

Tune to FM frequency: `GET /api/tune/set/93.3mhz`

Tune to AM frequency: `GET /api/tune/set/1040khz`

---

## Source Info

`GET /api/source/`

*Not yet implemented* 
