$(".post-delete").click(function (e) {
    e.preventDefault();
    var url = $(this).attr("href")
    fetch(url, {
        method: 'DELETE',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: ''
    }).then(data => {
        location.reload();
    });
})


