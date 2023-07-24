fetch('https://nofuss.co.za/names.php', {
    method: 'GET',
    headers: {
        'Accept': 'application/json',
    },
})
   .then(response => response.json())
   .then(response => {console.log(JSON.stringify(response)); document.write(JSON.stringify(response))})