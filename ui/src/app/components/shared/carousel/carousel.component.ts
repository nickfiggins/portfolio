import {
  Component,
  Input,
} from '@angular/core';

export interface ActiveSlides {
  previous: number;
  current: number;
  next: number;
}

@Component({
  selector: 'app-carousel',
  templateUrl: `carousel.component.html`,
  styleUrls: ['./carousel.component.scss'],
})
export class CarouselComponent {
  @Input()
  slides:any;
}
