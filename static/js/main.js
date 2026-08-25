newGameBtn = document.getElementById("new-game-button");
newGameForm = document.getElementById("new-game-form");
gameNameInput = document.getElementById("game-name-input");
confirmBtn = document.getElementById("create-game-confirm");
cancelBtn = document.getElementById("create-game-cancel");

fndGameBtn = document.getElementById("find-game-button");

newGameBtn.addEventListener("click", function () {
  fndGameBtn.style.display = "none"
  newGameBtn.style.display = "none";
  newGameForm.style.display = "block";
  gameNameInput.focus();
});

cancelBtn.addEventListener("click", function () {
  fndGameBtn.style.display = "block"
  newGameForm.style.display = "none";
  newGameBtn.style.display = "block";
  gameNameInput.value = "";
});

confirmBtn.addEventListener("click", async function () {
  const gameName = gameNameInput.value.trim();
  if (!gameName) {
    console.log("Enter the game name pls")
    return;
  }
  await CreateGame(gameName);
});


async function CreateGame(str, side) {
  try {
    const response = await fetch('http://localhost:9090/api/game/new', {
      method: 'POST',
      headers: {
        'Content-Type' : 'application/json'
      },
      credentials: 'include',
      body: JSON.stringify({
        name: str,
        role: side
      })
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error('Server side problem or Unexpected body')
    } else {
      console.log("all alright", data.body)
    }
  } catch (error) {
    console.error('Unexpected error', error);
  }
}
