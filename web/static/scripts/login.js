
var count = 0;

window.login = (event, form) => {
    event.preventDefault();

    if (document.body.classList.contains("loading")) {
        return;
    }

    document.body.classList.add("loading");

    const username = form.querySelector('#username').value;
    const password = form.querySelector('#password').value;

    fetch(`/api/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password }),
    }).then((response) => {
        return response.json();
    }).then((response) => {
        if (response.userType !== "none") {
            window.location.href = "/";
        } else {
            count += 1;
            document.body.classList.add("show-popup");
            document.body.classList.remove("loading");
            setTimeout(() => {
                count -= 1;
                if (count === 0) {
                    document.body.classList.remove("show-popup");
                }
            }, 3000);
        }
    }).catch(() => {
        document.body.classList.remove("loading");
    });
}
