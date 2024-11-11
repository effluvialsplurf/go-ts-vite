import ex from "excalibur";

const snakeGame = new ex.Engine({
    width: 1000,
    height: 1000,
    canvasElementId: "snakeGame",
    displayMode: ex.DisplayMode.FitScreen,
    pointerScope: ex.PointerScope.Document
});


const options = {
    pos: new ex.Vector(20, 20),
    width: 20,
    height: 20,
    color: ex.Color.Blue,
}


const snakeHead = new ex.Actor(options);
snakeHead.body.collisionType = ex.CollisionType.Active;

const snakeNode = new ex.Actor(options);
snakeNode.body.collisionType = ex.CollisionType.Active;

snakeHead.addChild(snakeNode);

snakeGame.add(snakeHead);

