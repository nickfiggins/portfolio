import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-freelance',
  templateUrl: './freelance.component.html',
  styleUrls: ['./freelance.component.scss']
})
export class FreelanceComponent implements OnInit {
  mobile: boolean = false;
  
  constructor() { }

  ngOnInit(): void {
  }

  onResize(event: any) {
    this.mobile = (event.target.innerWidth <= 400) ? true : false;
  }

}
