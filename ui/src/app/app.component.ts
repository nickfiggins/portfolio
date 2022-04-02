import { Component } from '@angular/core';
import { EnumService } from './components/shared/services/Enum.service';
import { IconService } from './components/shared/services/IconService';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent{
  title = 'portfolio';
  token:any;
  profile:any;
  particlesOptions = {};
  mobile:boolean = false;
  constructor(private iconService: IconService,
    private enums : EnumService) {
    this.particlesOptions = enums.particlesOptions;
  }

  ngOnInit() {
    this.iconService.registerIcons();
    if (window.screen.width === 360) { // 768px portrait
      this.mobile = true;
    }
  }

}
