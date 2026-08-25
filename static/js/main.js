newGameBtn = document.getElementById("new-game-button");

newGameBtn.addEventListener("click", function () {

  fetch('http://localhost:9090/api/game/new', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(test)
  })
    .then(response => response.json())
    .then(data => {
      console.log("Game created: ", data)
      const gameId = data.gameId;
    })
    .catch(error => {
      console.error('Error: ', error)
    });
})
