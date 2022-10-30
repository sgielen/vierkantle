export interface Board {
  width: number;
  height: number;
  cells: string[][];
}

export interface Coord {
  x: number;
  y: number;
}

export interface Path {
  coords: Coord[];
}
