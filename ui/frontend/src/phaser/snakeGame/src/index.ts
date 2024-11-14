import 'phaser';

const config: Phaser.Types.Core.GameConfig = {
    width: 600,
    height: 600,
    scale: {
        mode: Phaser.Scale.FIT,
        autoCenter: Phaser.Scale.CENTER_BOTH,
    },
    autoRound: false,
    parent: "game-container",
    physics: {
        default: "arcade",
        arcade: {
            debug: true,
        },
    },
    scene: [Game, GameOver],
};

export const game = new Phaser.Game(config);