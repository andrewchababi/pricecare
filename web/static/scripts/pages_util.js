function formatTimeComplex(timestamp) {
    let utc = new Date(timestamp);
    let now = new Date();
    if (utc.getFullYear() == now.getFullYear()
        && utc.getMonth() == now.getMonth()
        && utc.getDay() == now.getDay()) {
        let minutes = utc.getMinutes().toString().padStart(2, '0');
        return `${utc.getHours()}h${minutes}`
    } else {
        let minutes = utc.getMinutes().toString().padStart(2, '0');
        return `${utc.getHours()}h${minutes} - ${utc.getDay()}/${utc.getMonth()}/${utc.getFullYear()}`;
    }
}

function formatTimeSimple(timestamp) {
    let utc = new Date(timestamp);
    let now = new Date();

    if (utc.getFullYear() == now.getFullYear()
        && utc.getMonth() == now.getMonth()
        && utc.getDay() == now.getDay()) {
        let minutes = utc.getMinutes().toString().padStart(2, '0');
        return `${utc.getHours()}h${minutes}`
    } else {
        let minutes = utc.getMinutes().toString().padStart(2, '0');
        return `${utc.getDay()}/${utc.getMonth()}/${utc.getFullYear()}`;
    }
}

function generateUUID() {
    var d = new Date().getTime();
    var d2 = ((typeof performance !== 'undefined') && performance.now && (performance.now()*1000)) || 0;
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = Math.random() * 16;
        if(d > 0){
            r = (d + r)%16 | 0;
            d = Math.floor(d/16);
        } else {
            r = (d2 + r)%16 | 0;
            d2 = Math.floor(d2/16);
        }
        return (c === 'x' ? r : (r & 0x3 | 0x8)).toString(16);
    });
}
