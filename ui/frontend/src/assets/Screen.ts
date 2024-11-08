export default class Screen {
  public name: string;
  public height: number;
  public width: number;
  public canvas: HTMLCanvasElement;
  public ctx?: CanvasRenderingContext2D | null;

  CELL_SIZE = 20;

  public constructor(name: string, height: number, width: number) {
    this.name = name;
    this.height = height;
    this.width = width;
    this.canvas = <HTMLCanvasElement>document.createElement("canvas");
    this.createCanvas(this.canvas)
  }

  public createCanvas(canvas: HTMLCanvasElement): void {
    this.ctx = canvas.getContext("2d")!;
    this.ctx!.font = "25px serif";
    canvas.width = 500;
    canvas.height = 500;
    document.body.appendChild(canvas);
  }

  public drawScreen() {
    this.ctx?.beginPath();

    for (let x = 0; x < this.width + 1; x++) {
      this.ctx?.moveTo(this.CELL_SIZE * x, 0);
      this.ctx?.lineTo(this.CELL_SIZE * x, this.width * this.CELL_SIZE);
    }
    for (let y = 0; y < this.height + 1; y++) {
      this.ctx?.moveTo(0, this.CELL_SIZE * y);
      this.ctx?.lineTo(this.height * this.CELL_SIZE, this.CELL_SIZE * y);
    }

    this.ctx?.stroke();
  }
}
