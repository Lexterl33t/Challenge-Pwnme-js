/*
    YEAH !! Javascript obfuscation Challenge BAHAHAHHAHAHAAHHAHAHAHAHAHAHAHAHAHAHAH. 
*/

let get_uagent = (() => {
    return window.navigator.userAgent;
})

let get_timestamp = (() => {
    return Date.now()
})

let get_random = (() => {
    return Math.floor(Math.random() * 10)
})
let gen_axiome = ((timestamp) => {
    let randomNum = get_random()

    timestamp = (timestamp + 1337) % 1337
    console.log(timestamp)
    timestamp = timestamp << randomNum * 10
    console.log(randomNum)
    return timestamp
})


let gen_valid_token = ((uAgent, timestamp, key, axiome) => {

})

let init_hk = (() => {
    let userA = get_uagent()

    let time = get_timestamp()

    console.log(time)

    let axiome = gen_axiome(time)

    console.log(axiome)
    /*if (userA != "l33t_Akeur") 
        return*/
    
    
    
})
