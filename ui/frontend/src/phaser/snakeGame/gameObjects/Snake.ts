class SnakeNode extends Phaser.GameObjects.Rectangle {
    declare node: Phaser.Physics.Arcade.Body;
    declare isHead: boolean;

    constructor(scene: Phaser.Scene, x: number, y: number, isHead: boolean, number?: number) {
        super(scene, x, y, 32, 32, 0x9600ff);
        this.setOrigin(0.5);
        this.scene.add.existing(this);
        this.scene.physics.add.existing(this);
        this.setScale(1);
    }
}

export default SnakeNode;