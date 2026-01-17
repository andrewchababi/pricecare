
window.onpageshow = (event) => {
    if (event.persisted) {
        window.location.reload();
    }
};
