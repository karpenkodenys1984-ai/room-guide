export namespace dto {
	
	export class RoomData {
	    Label: string;
	    Type: string;
	
	    static createFrom(source: any = {}) {
	        return new RoomData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Label = source["Label"];
	        this.Type = source["Type"];
	    }
	}
	export class Position {
	    X: number;
	    Y: number;
	
	    static createFrom(source: any = {}) {
	        return new Position(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.X = source["X"];
	        this.Y = source["Y"];
	    }
	}
	export class Node {
	    Id: number;
	    Type: string;
	    Position: Position;
	    Data: RoomData;
	    Draggable?: boolean;
	    Selectable?: boolean;
	    Connectable?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Node(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Type = source["Type"];
	        this.Position = this.convertValues(source["Position"], Position);
	        this.Data = this.convertValues(source["Data"], RoomData);
	        this.Draggable = source["Draggable"];
	        this.Selectable = source["Selectable"];
	        this.Connectable = source["Connectable"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

export namespace logger {
	
	export class FrontendLogEntry {
	    level: string;
	    message: string;
	    component: string;
	    extra?: Record<string, any>;
	    timestamp: string;
	
	    static createFrom(source: any = {}) {
	        return new FrontendLogEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.level = source["level"];
	        this.message = source["message"];
	        this.component = source["component"];
	        this.extra = source["extra"];
	        this.timestamp = source["timestamp"];
	    }
	}

}

