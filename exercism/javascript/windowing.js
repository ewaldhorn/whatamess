// @ts-check

export function Size(width,height) {
  this.width=width ?? 80;
  this.height=height ?? 60;
};

Size.prototype.resize = function(newWidth,newHeight) {
  this.width=newWidth;
  this.height=newHeight;
};

export function Position(x,y) {
  this.x=x ?? 0;
  this.y=y ?? 0;
};

Position.prototype.move = function(newX,newY) {
  this.x=newX;
  this.y=newY;
};

export class ProgramWindow {
  constructor() {
    this.screenSize = new Size(800,600);
    this.size = new Size();
    this.position = new Position();
  }

  resize(newSize) {
    // sanity check input
    let w = Math.max(newSize.width, 1);
    let h = Math.max(newSize.height, 1);

    // how much space is available
    let wRemain = 800 - this.position.x;
    let hRemain = 600 - this.position.y;

    if (wRemain < w) {
      w = wRemain;
    }

    if (hRemain < h) {
      h = hRemain;
    }

    this.size = new Size(w,h);
  }

  move(newPosition) {
    // sanity check input
    let x = Math.max(newPosition.x, 0);
    let y = Math.max(newPosition.y, 0);

    // will this fit at the new position
    let wRemain = 800 - (newPosition.x + this.size.width);
    let hRemain = 600 - (newPosition.y + this.size.height);

    if (wRemain < 0) {
      x = x + wRemain;
    }

    if (hRemain < 0) {
      y = y + hRemain;
    }

    this.position = new Position(x,y);
  }
}

export function changeWindow(newWindow) {
  newWindow.resize(new Size(400,300));
  newWindow.move(new Position(100,150));
  return newWindow;
}
