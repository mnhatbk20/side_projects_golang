document.querySelector(".ajax-logout").addEventListener("click", (e) => {
    e.preventDefault();
    const response = fetch('/api/logout', {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: ''
    }).then(data => {
        location.reload();
    });
});