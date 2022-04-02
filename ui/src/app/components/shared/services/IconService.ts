import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";
import { Injectable } from "@angular/core";
import { EnumService } from "./Enum.service";


@Injectable({
    providedIn: 'root'
  })
  export class IconService {
  
    constructor(
      private matIconRegistry: MatIconRegistry,
      private domSanitizer: DomSanitizer,
      private Enum: EnumService
    ) { }
  
    public registerIcons(): void {
      this.loadIcons(Object.values(this.Enum.Icons), '../assets/svg');
    }
  
    private loadIcons(iconKeys: string[], iconUrl: string): void {
      iconKeys.forEach(key => {
        this.matIconRegistry.addSvgIcon(key, this.domSanitizer.bypassSecurityTrustResourceUrl(`${iconUrl}/${key}.svg`));
      });
    }
  }
  