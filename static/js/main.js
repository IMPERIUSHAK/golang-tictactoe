newGameBtn = document.getElementById("new-game-button");
newGameForm = document.getElementById("new-game-form");
gameNameInput = document.getElementById("game-name-input");
confirmBtn = document.getElementById("create-game-confirm");
cancelBtn = document.getElementById("create-game-cancel");

fndGameBtn = document.getElementById("find-room-button");
fndGameForm = document.getElementById("find-room-form")
jconfirmBtn = document.getElementById("join-confirm-button")
jcancelBtn = document.getElementById("join-cancel-button")
roomIdInput = document.getElementById("room-id-input")

xrole = document.getElementById("xrole")
orole = document.getElementById("orole")

newGameBtn.addEventListener("click", function () {
  fndGameBtn.style.display = "none"
  newGameBtn.style.display = "none";

  newGameForm.style.display = "block";
  gameNameInput.focus();
});

cancelBtn.addEventListener("click", function () {
  fndGameBtn.style.display = "block"
  newGameBtn.style.display = "block";

  newGameForm.style.display = "none";
  gameNameInput.value = "";
});

confirmBtn.addEventListener("click", async function () {
  const gameName = gameNameInput.value.trim();
  let role;
  if (!gameName) {
    console.log("Enter the game name pls")
    return;
  }

  if (xrole.checked) {
    role = "x"
  } else {
    role = "o"
  }

  await CreateGame(gameName, role);
});

async function CreateGame(str, side) {
  try {
    const response = await fetch('api/game/new', {
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
      console.log("Succes", data)
    }
  } catch (error) {
    console.error('Unexpected error', error);
  }
}

fndGameBtn.addEventListener("click", function () {

  fndGameBtn.style.display = "none";
  newGameBtn.style.display = "none";

  fndGameForm.style.display = "block"
  roomIdInput.focus();
});

jcancelBtn.addEventListener("click", function () {

  fndGameBtn.style.display = "block";
  newGameBtn.style.display = "block";

  fndGameForm.style.display = "none"
  roomIdInput.value = "";
});

jconfirmBtn.addEventListener("click", async function () {

  const roomId = roomIdInput.value.trim();
  if (!roomId) {
    console.log("Enter room Id please");
    return;
  }

  await JoinRoom(roomId);
});

async function JoinRoom(roomId) {
  const url = `http://localhost:9090/api/game/join/${roomId}`;

  try {
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include'
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error('Server side problem or Unexpected body')
    } else {
      console.log("Succes", data)
    }

  } catch (error) {
    console.error('Unexpected error', error);
  }
}
