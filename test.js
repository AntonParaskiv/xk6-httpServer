import { sleep } from 'k6';

import Factory from 'k6/x/httpServer';

export default function () {
    let api = Factory.new()
        .get('/custom', (req, res) => {
            console.log(req.body)
            let response = {foo:"bar"}
            res.json(response)
        })
        .post('/mirror', (req, res) => {
            console.log(req.body)
            res.json(req.body)
        })
        .start() // random port
        // or .start(1111) for specific port

    console.log(`http://${api.addr()}`)

    sleep(60)
    api.stop()
}